package operate

// desc: 订单操作

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	pb_timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"github.com/syndtr/goleveldb/leveldb/util"
	pb_item "github.com/zhanglp92/rep/api/pb/item"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/db"
	"go.uber.org/zap"
)

// Form ...
type Form struct {
	config *config.Config

	user *User

	logger *zap.Logger

	// mu sync.RWMutex

	sizeReg *regexp.Regexp
}

func newForm(config *config.Config, user *User) (a *Form, err error) {

	a = &Form{config: config, user: user}
	a.logger = config.Logger().Named("op-form")
	if a.sizeReg, err = regexp.Compile(`[\d\.]+[mM]+`); err != nil {
		return nil, err
	}

	return a, a.init()
}

func (a *Form) init() error { return nil }

func (a *Form) idToKey(id int32) []byte {
	return []byte(db.PKForm + strconv.FormatInt(int64(id), 10))
}

// Get ...
func (a *Form) Get(id int32) (item *pb_item.Item, err error) {
	if item, err = a.get(id); err != nil {
		return nil, err
	}

	if item.User, err = a.user.Get(item.GetPhone()); err != nil {
		a.logger.Warn("get user fail", zap.String("userid", item.GetPhone()), zap.Error(err))
	}
	return
}

func (a *Form) genID() (id int32) {
	for _, item := range a.Range() {
		if id < item.GetId() {
			id = item.GetId()
		}
	}
	return id + 1
}

// Put ...
func (a *Form) Put(param *param) error {
	item := param.toForm()

	if item.GetId() <= 0 {
		item.Id = a.genID()
	}

	if item.GetId() <= 0 {
		return fmt.Errorf("用户名为空")
	}

	user, err := a.user.Get(item.GetPhone())
	if err != nil {
		return fmt.Errorf("nomatch user[%v]", item.GetPhone())
	}
	item.Username = user.GetName()

	old, err := a.get(item.GetId())
	if err != nil && err != errors.ErrNotFound {
		return err
	}

	body, err := proto.Marshal(a.itemAdjust(a.mergeItem(old, item)))
	if err != nil {
		return err
	}
	return db.DB().Put(a.idToKey(item.GetId()), body, nil)
}

// Del ...
func (a *Form) Del(id int32) error {
	return db.DB().Delete(a.idToKey(id), nil)
}

// Range ...
func (a *Form) Range() (dq []*pb_item.Item) {
	dq = make([]*pb_item.Item, 0)

	it := db.DB().NewIterator(&util.Range{Start: []byte(db.PKForm), Limit: []byte(db.PKForm + "Z")}, nil)
	for it.Next() {
		if it.Valid() {

			var (
				item pb_item.Item
				err  error
			)

			if err := proto.Unmarshal(it.Value(), &item); err != nil {
				continue
			}

			if item.User, err = a.user.Get(item.GetPhone()); err != nil {
				a.logger.Warn("get user fail", zap.String("phone", item.GetPhone()), zap.Error(err))
			}
			item.Username = item.User.GetName()

			dq = append(dq, a.itemAdjust(&item))
		}
	}
	return dq
}

// ---- internal ----

func (a *Form) get(id int32) (*pb_item.Item, error) {
	body, err := db.DB().Get(a.idToKey(id), nil)
	if err != nil {
		return nil, err
	}

	var item pb_item.Item
	if err := proto.Unmarshal(body, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// t left join s
func (a *Form) mergeItem(t, s *pb_item.Item) (item *pb_item.Item) {
	defer func() {
		a.itemAdjust(item)

		now := time.Now()
		item.UpdateTime = &pb_timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		}
	}()

	if s == nil {
		return t
	}
	if t == nil {
		return s
	}

	vt := reflect.ValueOf(*t)
	vs := reflect.ValueOf(*s)
	vtElem := reflect.ValueOf(t).Elem()

	for i := 0; i < vt.NumField(); i++ {
		fs := vs.Field(i)
		if !reflect.DeepEqual(fs.Interface(), reflect.Zero(fs.Type()).Interface()) {
			if !vtElem.Field(i).CanSet() {
				continue
			}
			vtElem.Field(i).Set(fs)
		}
	}
	return t
}

func (a *Form) itemAdjust(item *pb_item.Item) *pb_item.Item {
	if item == nil {
		return nil
	}

	//  计算价格
	a.calPrice(item)

	if item.GetCreateTime() == nil {
		now := time.Now()
		item.CreateTime = &pb_timestamp.Timestamp{
			Seconds: now.Unix(),
			Nanos:   int32(now.Nanosecond()),
		}
	}
	if item.GetUpdateTime() == nil {
		item.UpdateTime = &pb_timestamp.Timestamp{
			Seconds: item.GetCreateTime().GetSeconds(),
			Nanos:   item.GetCreateTime().GetNanos(),
		}
	}

	{ // 调整显示时间
		t := item.GetCreateTime()
		item.Screate = time.Unix(t.GetSeconds(), int64(t.GetNanos())).Format("2006/01/02 15:04:05")

		t = item.GetUpdateTime()
		item.Supdate = time.Unix(t.GetSeconds(), int64(t.GetNanos())).Format("2006/01/02 15:04:05")
	}

	return item
}

func (a *Form) calPrice(item *pb_item.Item) {
	// 按套计算总价
	if item.GetPriceset() > 0 {
		item.Price = item.GetPriceset()
		return
	}

	if n, ok := a.sqCkg(item); ok {
		item.Squaremeter = n
	} else if n, ok := a.sqDirect(item); ok {
		item.Squaremeter = n
	}

	// 平方米 * 单价 = 总价
	if item.GetPriceunit() <= 0 {
		return
	}

	item.Price = item.GetSquaremeter() * item.GetPriceunit()
}

// 直接提取(平方米)
func (a *Form) sqDirect(item *pb_item.Item) (v float32, ok bool) {
	res := a.sizeReg.FindAll([]byte(item.GetSize()), -1)
	if len(res) <= 0 {
		return
	}

	src := strings.TrimSpace(strings.ToUpper(string(res[len(res)-1])))
	if strings.HasSuffix(src, "MM") {
		n, err := strconv.ParseFloat(src[:len(src)-2], 0)
		if err != nil {
			return
		}
		return float32(n / 1000), true
	}

	if strings.HasSuffix(src, "M") {
		n, err := strconv.ParseFloat(src[:len(src)-1], 0)
		if err != nil {
			return
		}
		return float32(n), true
	}

	return 0, false
}

// 通过长宽高计算面积(长*宽=平方米)
func (a *Form) sqCkg(item *pb_item.Item) (v float32, ok bool) {
	segs := strings.Split(item.GetSize(), "*")
	if len(segs) < 2 {
		return
	}

	c, err := strconv.ParseFloat(strings.TrimSpace(segs[0]), 0)
	if err != nil {
		return
	}

	k, err := strconv.ParseFloat(strings.TrimSpace(segs[1]), 0)
	if err != nil {
		return
	}

	return float32(c / 1000 * k / 1000), true
}

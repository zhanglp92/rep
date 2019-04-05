package operate

// desc: 订单操作

import (
	"fmt"
	"reflect"
	"strconv"

	"go.uber.org/zap"

	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb/util"
	pb_item "github.com/zhanglp92/rep/api/pb/item"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/db"
)

// Form ...
type Form struct {
	config *config.Config

	user *User

	logger *zap.Logger
}

func newForm(config *config.Config, user *User) (*Form, error) {
	a := &Form{config: config, user: user}
	a.logger = config.Logger().Named("op-form")

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

	if item.User, err = a.user.Get(item.GetUserid()); err != nil {
		a.logger.Warn("get user fail", zap.Int32("userid", item.GetUserid()), zap.Error(err))
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
func (a *Form) Put(item *pb_item.Item) error {
	if item.GetId() <= 0 {
		return ErrBadData
	}

	old, err := a.get(item.GetId())
	if err != nil && err != ErrFormNotExists {
		return err
	}

	uid := item.GetUser().GetId()
	item.User = nil
	item = a.mergeItem(old, item)
	if item.Userid <= 0 {
		item.Userid = uid
	}

	if _, err := a.user.Get(item.GetUserid()); err != nil {
		return fmt.Errorf("nomatch user[%v]", item.GetUserid())
	}

	body, err := proto.Marshal(item)
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

			if item.User, err = a.user.Get(item.GetUserid()); err != nil {
				a.logger.Warn("get user fail", zap.Int32("userid", item.GetUserid()), zap.Error(err))
			}

			dq = append(dq, &item)
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
func (a *Form) mergeItem(t, s *pb_item.Item) *pb_item.Item {
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

package operate

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/syndtr/goleveldb/leveldb/util"
	pb_user "github.com/zhanglp92/rep/api/pb/user"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/db"
)

// desc: 用户相关

// User ...
type User struct {
	config *config.Config
}

func newUser(config *config.Config) (*User, error) {
	a := &User{config: config}
	return a, a.init()
}

func (a *User) init() error { return nil }

func (a *User) idToKey(id int32) []byte {
	return []byte(db.PKUser + strconv.FormatInt(int64(id), 10))
}

// Get ...
func (a *User) Get(id int32) (*pb_user.Item, error) {
	if id <= 0 {
		return nil, ErrBadData
	}

	body, err := db.DB().Get(a.idToKey(id), nil)
	if err != nil {
		return nil, err
	}

	var item pb_user.Item
	if err := proto.Unmarshal(body, &item); err != nil {
		return nil, err
	}
	return &item, nil
}

// 生成新id 并check
func (a *User) genID(tar *pb_user.Item) (id int32, isExist bool) {
	for _, item := range a.Range() {
		if strings.Compare(item.GetName(), tar.GetName()) == 0 &&
			strings.Compare(item.GetPhone(), tar.GetPhone()) == 0 {
			return item.GetId(), true
		}

		if id < item.GetId() {
			id = item.GetId()
		}
	}
	return id + 1, false
}

// Put ...
func (a *User) Put(param *param) (err error) {
	var item pb_user.Item
	if err = param.parseObj(&item); err != nil {
		return err
	}

	var isExist bool
	if item.Id, isExist = a.genID(&item); isExist && !param.isover {
		return fmt.Errorf("user[%v-%v] is exists", item.GetName(), item.GetPhone())
	}

	body, err := proto.Marshal(&item)
	if err != nil {
		return err
	}
	return db.DB().Put(a.idToKey(item.GetId()), body, nil)
}

// Del ...
func (a *User) Del(id int32) error {
	return db.DB().Delete(a.idToKey(id), nil)
}

// Range ...
func (a *User) Range() (dq []*pb_user.Item) {
	it := db.DB().NewIterator(&util.Range{Start: []byte(db.PKUser), Limit: []byte(db.PKUser + "Z")}, nil)
	for it.Next() {
		if it.Valid() {
			var item pb_user.Item
			if err := proto.Unmarshal(it.Value(), &item); err != nil {
				continue
			}
			dq = append(dq, &item)
		}
	}
	return dq
}

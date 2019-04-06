package operate

import (
	"fmt"
	"time"

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

func (a *User) idToKey(id string) []byte {
	return []byte(db.PKUser + id)
}

// Get ...
func (a *User) Get(id string) (*pb_user.Item, error) {
	if len(id) <= 0 {
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
func (a *User) genID(tar *pb_user.Item) (id string, isExist bool) {
	id = tar.GetPhone()
	if v, err := a.Get(id); v != nil && err == nil {
		isExist = true
	}
	return id, isExist
}

// Put ...
func (a *User) Put(param *param) (err error) {
	var (
		isExist bool
		item    = param.toUser()
	)

	if item.Id, isExist = a.genID(item); isExist && !param.isover {
		return fmt.Errorf("user[%v-%v] is exists", item.GetName(), item.GetPhone())
	}

	body, err := proto.Marshal(item)
	if err != nil {
		return err
	}
	return db.DB().Put(a.idToKey(item.GetId()), body, nil)
}

// Del ...
func (a *User) Del(id string) error {
	return db.DB().Delete(a.idToKey(id), nil)
}

// Range ...
func (a *User) Range() (dq []*pb_user.Item) {
	dq = make([]*pb_user.Item, 0)

	it := db.DB().NewIterator(&util.Range{Start: []byte(db.PKUser), Limit: []byte(db.PKUser + "Z")}, nil)
	for it.Next() {
		if it.Valid() {
			var item pb_user.Item
			if err := proto.Unmarshal(it.Value(), &item); err != nil {
				continue
			}
			dq = append(dq, a.itemAdjust(&item))
		}
	}
	return dq
}

func (a *User) itemAdjust(item *pb_user.Item) *pb_user.Item {
	if item == nil {
		return nil
	}

	if len(item.GetScreate()) <= 0 {
		t := item.GetCreateTime()
		item.Screate = time.Unix(t.GetSeconds(), int64(t.GetNanos())).Format("2006/01/02 15:04:05")
	}

	return item
}

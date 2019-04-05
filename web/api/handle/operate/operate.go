package operate

// desc: 订单

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	dbe "github.com/syndtr/goleveldb/leveldb/errors"
	pb_item "github.com/zhanglp92/rep/api/pb/item"
	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web/api/handle"
)

// Err ...
var (
	ErrFormNotExists = dbe.ErrNotFound
	ErrBadData       = errors.New("bad data")
)

func init() {
	handle.Register(&Operate{location: "/op"})
}

// Operate ...
type Operate struct {
	*bb.Base

	config *config.Config

	location string

	form *Form

	user *User
}

// Init ...
func (a *Operate) Init(config *config.Config) (err error) {
	a.config = config

	if a.Base, err = bb.New(bb.CurrentBaseName(), a.config.Logger()); err != nil {
		return err
	}

	if a.user, err = newUser(a.config); err != nil {
		return err
	}

	if a.form, err = newForm(a.config, a.user); err != nil {
		return err
	}

	return
}

// ServerHTTP ...
func (a *Operate) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var body []byte
	param, err := newParam(r)

	fmt.Println("TT: ", param.String(), err)

	if err != nil {
		goto END
	}

	switch param.ty {
	case "form":
		body, err = a.opForm(param)
	case "user":
		body, err = a.opUser(param)
	default:
		err = fmt.Errorf("type err, candidate[form, user]")
	}

END:
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write(body)
	}
}

// Location ...
func (a *Operate) Location() string {
	return a.location
}

// ---- internal ----

func (a *Operate) opUser(param *param) ([]byte, error) {
	var (
		obj interface{}
		err error
	)

	switch param.op {
	case "get":
		obj, err = a.user.Get(param.id)
	case "put":
		err = a.user.Put(param)
	case "del":
		err = a.user.Del(param.id)
	case "range":
		obj = a.user.Range()
	default:
		err = fmt.Errorf("op err, candidate[get, put, del, range]")
	}

	if err != nil {
		return nil, err
	}

	return json.Marshal(obj)
}

func (a *Operate) opForm(param *param) ([]byte, error) {
	var (
		obj interface{}
		err error
	)

	switch param.op {
	case "get":
		obj, err = a.form.Get(param.id)
	case "put":
		var item pb_item.Item
		if err = param.parseObj(&item); err == nil {
			err = a.form.Put(&item)
		}
	case "del":
		err = a.form.Del(param.id)
	case "range":
		obj = a.form.Range()
	default:
		err = fmt.Errorf("op err, candidate[get, put, del, range]")
	}

	if err != nil {
		return nil, err
	}

	return json.Marshal(obj)
}

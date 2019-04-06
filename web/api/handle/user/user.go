package form

// desc: 订单

import (
	"html/template"
	"net/http"
	"path/filepath"

	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web/api/handle"
	"github.com/zhanglp92/rep/web/api/handle/base"
)

func init() {
	handle.Register(&User{location: "/user"})
}

// User ...
type User struct {
	*bb.Base

	config *config.Config

	location string
}

// Init ...
func (a *User) Init(config *config.Config) (err error) {
	a.config = config

	if a.Base, err = bb.New(bb.CurrentBaseName(), a.config.Logger()); err != nil {
		return err
	}
	return
}

// ServerHTTP ...
func (a *User) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	switch r.Method {
	case http.MethodGet:
		err = a.apiGet(w, r)
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			err = a.apiPost(w, r)
		}
	default:
		err = base.ErrMethod
	}

	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

// Location ...
func (a *User) Location() string {
	return a.location
}

// ---- internal ----

func (a *User) apiGet(w http.ResponseWriter, r *http.Request) error {
	path := filepath.Join(bb.CurrentDir(), "html/index.html")

	t := template.Must(template.ParseFiles(path))
	return t.Execute(w, nil)
}

func (a *User) apiPost(w http.ResponseWriter, r *http.Request) error {
	return nil
}

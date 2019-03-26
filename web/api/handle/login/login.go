package login

// desc: 登录

import (
	"html/template"
	"net/http"
	"path/filepath"

	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web/api/handle"
)

func init() {
	handle.Register(&Login{location: "/login"})
}

// Login ...
type Login struct {
	*bb.Base

	config *config.Config

	location string
}

// Init ...
func (a *Login) Init(config *config.Config) (err error) {
	a.config = config

	if a.Base, err = bb.New(bb.CurrentBaseName(), a.config.Logger()); err != nil {
		return err
	}
	return
}

// ServerHTTP ...
func (a *Login) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(bb.CurrentDir(), "html/login.gtpl")

	t, err := template.ParseFiles(path)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		if err := t.Execute(w, nil); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

// Location ...
func (a *Login) Location() string {
	return a.location
}

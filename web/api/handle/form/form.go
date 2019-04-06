package form

// desc: 订单

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web/api/handle"
	"github.com/zhanglp92/rep/web/api/handle/base"
)

func init() {
	handle.Register(&Form{location: "/form"})
}

// Form ...
type Form struct {
	*bb.Base

	config *config.Config

	location string
}

// Init ...
func (a *Form) Init(config *config.Config) (err error) {
	a.config = config

	if a.Base, err = bb.New(bb.CurrentBaseName(), a.config.Logger()); err != nil {
		return err
	}
	return
}

// ServerHTTP ...
func (a *Form) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	fmt.Println("TTxx: ", r.PostForm)
	fmt.Println("TTx1: ", r.Form)

	switch r.Method {
	case http.MethodGet:
		a.apiGet(w, r)
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			a.apiPost(w, r)
		}
	default:
		err = base.ErrMethod
	}

	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

// Location ...
func (a *Form) Location() string {
	return a.location
}

// ---- internal ----

func (a *Form) apiGet(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join(bb.CurrentDir(), "html/index.html")

	t, err := template.ParseFiles(path)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		if err := t.Execute(w, nil); err != nil {
			w.Write([]byte(err.Error()))
		}
	}
}

func (a *Form) apiPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("username:", r.Form["username"])
	fmt.Println("password:", r.Form["password"])
}

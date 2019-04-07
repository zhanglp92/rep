package form

// desc: 订单

import (
	"html/template"
	"net/http"

	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web/api/handle"
	"github.com/zhanglp92/rep/web/api/handle/base"
)

func init() {
	handle.Register(&Display{location: "/display"})
}

// Display ...
type Display struct {
	*bb.Base

	config *config.Config

	location string
}

// Init ...
func (a *Display) Init(config *config.Config) (err error) {
	a.config = config

	if a.Base, err = bb.New(bb.CurrentBaseName(), a.config.Logger()); err != nil {
		return err
	}
	return
}

// ServerHTTP ...
func (a *Display) ServerHTTP(w http.ResponseWriter, r *http.Request) {
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
func (a *Display) Location() string {
	return a.location
}

// ---- internal ----

func (a *Display) apiGet(w http.ResponseWriter, r *http.Request) error {
	path := "lib/html/display/index.html"

	t := template.Must(template.ParseFiles(path))
	return t.Execute(w, nil)
}

func (a *Display) apiPost(w http.ResponseWriter, r *http.Request) error {
	return nil
}

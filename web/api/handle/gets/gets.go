package form

// desc: 订单

import (
	"html/template"
	"log"
	"net/http"

	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web/api/handle"
	"github.com/zhanglp92/rep/web/api/handle/base"
)

func init() {
	handle.Register(&Gets{location: "/gets"})
}

// Gets ...
type Gets struct {
	*bb.Base

	config *config.Config

	location string
}

// Init ...
func (a *Gets) Init(config *config.Config) (err error) {
	a.config = config

	if a.Base, err = bb.New(bb.CurrentBaseName(), a.config.Logger()); err != nil {
		return err
	}
	return
}

// ServerHTTP ...
func (a *Gets) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var err error

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
func (a *Gets) Location() string {
	return a.location
}

// ---- internal ----

func (a *Gets) apiGet(w http.ResponseWriter, r *http.Request) {
	a.config.DataPath()

	t := template.Must(template.ParseFiles(path))
	names := []string{"john", "jim"}
	if err := t.Execute(w, names); err != nil {
		log.Fatal(err)
	}
}

func (a *Gets) apiPost(w http.ResponseWriter, r *http.Request) {
}

package webcome

// desc: 欢迎

import (
	"html/template"
	"net/http"
	"path/filepath"

	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web/api/handle"
	"go.uber.org/zap"
)

func init() {
	handle.Register(&Welcome{location: "/welcome"})
}

// Welcome ...
type Welcome struct {
	*bb.Base

	config *config.Config

	location string
}

// Init ...
func (a *Welcome) Init(config *config.Config) (err error) {
	a.config = config

	if a.Base, err = bb.New(bb.CurrentBaseName(), a.config.Logger()); err != nil {
		return err
	}
	return
}

// ServerHTTP ...
func (a *Welcome) ServerHTTP(w http.ResponseWriter, r *http.Request) {

	path := filepath.Join(bb.CurrentDir(), "html/welcome.gtpl")
	a.Logger().Info("html", zap.String("path", path))

	t, err := template.ParseFiles(path)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err := t.Execute(w, nil); err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	return
}

// Location ...
func (a *Welcome) Location() string {
	return a.location
}

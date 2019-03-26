package handle

import (
	"net/http"

	"github.com/zhanglp92/rep/config"
)

var routers []Handle

// Routers ...
func Routers() []Handle {
	return routers
}

// Init ...
func Init(config *config.Config) (err error) {
	for _, h := range Routers() {
		if err = h.Init(config); err != nil {
			return err
		}
	}
	return
}

// Handle ...
type Handle interface {
	Init(*config.Config) error

	ServerHTTP(http.ResponseWriter, *http.Request)

	Location() string
}

// Register ...
func Register(h Handle) {
	routers = append(routers, h)
}

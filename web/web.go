package web

// desc: web 接口

import (
	"fmt"
	"net"
	"net/http"
	"sync/atomic"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	v1 "github.com/zhanglp92/rep/web/api/v1"
	route "github.com/zhanglp92/rep/web/router"
	"go.uber.org/zap"
	"golang.org/x/net/netutil"
)

// Web ...
type Web struct {
	*bb.Base

	listenErrCh chan error

	ready uint32

	config *config.Config

	apiV1 *v1.API

	router *route.Router
}

// New ...
func New(config *config.Config) (w *Web, err error) {
	w = &Web{
		config:      config,
		listenErrCh: make(chan error, 1),
	}

	if w.Base, err = bb.New(bb.CurrentBaseName(), w.config.Logger()); err != nil {
		return nil, err
	}

	return w, w.init()
}

// Run serves the HTTP endpoints.
func (a *Web) Run() {
	operationName := nethttp.OperationNameFunc(func(r *http.Request) string {
		return fmt.Sprintf("%s %s", r.Method, r.URL.Path)
	})

	server := &http.Server{
		Addr:        a.config.Web().Addr,
		Handler:     nethttp.Middleware(opentracing.GlobalTracer(), a.router, operationName),
		ReadTimeout: a.config.Web().ReadTimeOut,
	}

	a.Logger().Info("Bind", zap.String("addr", server.Addr))

	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		a.listenErrCh <- err
	} else {
		limitedListener := netutil.LimitListener(listener, a.config.Web().MaxConnections)
		a.listenErrCh <- server.Serve(limitedListener)
	}
}

// ListenError returns the receive-only channel that signals errors while starting the web server.
func (a *Web) ListenError() <-chan error {
	return a.listenErrCh
}

// ---- internal ----

func (a *Web) init() (err error) {
	if a.apiV1, err = v1.New(a.config); err != nil {
		return
	}

	return a.initRouter()
}

func (a *Web) initRouter() (err error) {
	router := route.New()
	a.router = router

	wc := a.config.Web()

	if wc.RoutePrefix != "/" {
		// If the prefix is missing for the root path, prepend it.
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, wc.RoutePrefix, http.StatusFound)
		})
		router = router.WithPrefix(wc.RoutePrefix)
	}

	readyf := a.testReady

	a.apiV1.Register(router.WithPrefix("/api/v1"))

	router.Get("/-/ready", readyf(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Dmp is Ready.\n")
	}))

	return nil
}

// Verifies whether the server is ready or not.
func (a *Web) isReady() bool {
	return atomic.LoadUint32(&a.ready) != 0
}

// Checks if server is ready, calls f if it is, returns 503 if it is not.
func (a *Web) testReady(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if a.isReady() {
			f(w, r)
		} else {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintf(w, "Service Unavailable")
		}
	}
}

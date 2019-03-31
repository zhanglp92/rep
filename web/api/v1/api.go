package v1

// desc: 接口实例

import (
	"encoding/json"
	"net/http"

	bb "github.com/zhanglp92/rep/base"
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web/api/handle"
	route "github.com/zhanglp92/rep/web/router"
	"go.uber.org/zap"
)

type status string

const (
	statusSuccess status = "success"
	statusError   status = "error"
)

type errorType string

const (
	errorTimeout  = "timeout"
	errorCanceled = "canceled"
	errorExec     = "execution"
	errorBadPara  = "bad_para"
	errorBadData  = "bad_data"
	errorInternal = "internal"
)

var corsHeaders = map[string]string{
	"Access-Control-Allow-Headers":  "Accept, Authorization, Content-Type, Origin",
	"Access-Control-Allow-Methods":  "GET, OPTIONS",
	"Access-Control-Allow-Origin":   "*",
	"Access-Control-Expose-Headers": "Date",
}

type apiError struct {
	typ errorType
	err error
}

type response struct {
	Status    status      `json:"status"`
	Data      interface{} `json:"data,omitempty"`
	ErrorType errorType   `json:"errorType,omitempty"`
	Error     string      `json:"error,omitempty"`
}

// Enables cross-site script calls.
func setCORS(w http.ResponseWriter) {
	for h, v := range corsHeaders {
		w.Header().Set(h, v)
	}
}

type apiFunc func(r *http.Request) (interface{}, *apiError)

// API ...
type API struct {
	*bb.Base

	config *config.Config
}

// New ...
func New(config *config.Config) (a *API, err error) {
	a = &API{config: config}

	if a.Base, err = bb.New(bb.CurrentBaseName(), a.config.Logger()); err != nil {
		return nil, err
	}

	if err := handle.Init(config); err != nil {
		return nil, err
	}

	return a, a.init()
}

// Register the API's endpoints in the given router.
func (a *API) Register(r *route.Router) {
	instr := func(name string, f apiFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			setCORS(w)
			if data, err := f(r); err != nil {
				respondError(w, err, data)
			} else if data != nil {
				respond(w, data)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
		})
	}
	_ = instr

	for _, h := range handle.Routers() {
		a.Logger().Info("register", zap.String("location", h.Location()))

		r.Get(h.Location(), h.ServerHTTP)
		r.Post(h.Location(), h.ServerHTTP)
	}
}

// ---- internal ----

func (a *API) init() error {
	return nil
}

func respondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	b, err := json.Marshal(data)
	if err != nil {
		return
	}
	w.Write(b)
}

func respondError(w http.ResponseWriter, apiErr *apiError, data interface{}) {
	var code int
	switch apiErr.typ {
	case errorBadData:
		code = http.StatusBadRequest
	case errorExec:
		code = 422
	case errorCanceled, errorTimeout:
		code = http.StatusServiceUnavailable
	case errorInternal:
		code = http.StatusInternalServerError
	default:
		code = http.StatusInternalServerError
	}

	resp := &response{
		Status:    statusError,
		ErrorType: apiErr.typ,
		Error:     apiErr.err.Error(),
		Data:      data,
	}

	respondJSON(w, code, resp)
}

func respond(w http.ResponseWriter, data interface{}) {
	resp := &response{
		Status: statusSuccess,
		Data:   data,
	}

	respondJSON(w, http.StatusOK, resp)
}

package config

import (
	"net/url"
	"time"
)

// desc: web 配置

// Web ...
type Web struct {
	Addr           string
	ReadTimeOut    time.Duration
	WriteTimeOut   time.Duration
	MaxConnections int
	ExternalURLSrc string
	RoutePrefix    string
	ExternalURL    *url.URL
}

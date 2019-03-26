package config

// desc: 全局配置

import (
	bb "github.com/zhanglp92/rep/base"
)

// Config ...
type Config struct {
	*bb.Base

	path string

	// 内部配置
	cfg *cfg
}

// New ...
func New(path string) (c *Config, err error) {
	c = &Config{path: path}

	// 加载配置
	if c.cfg, err = newCfg(path); err != nil {
		return nil, err
	}

	// 初始化logger
	if c.Base, err = bb.New(bb.CurrentBaseName(), newLogger(&c.cfg.Logger)); err != nil {
		return nil, err
	}

	return c, c.init()
}

// Web config
func (a *Config) Web() *Web {
	if a == nil {
		return nil
	}
	return &a.cfg.Web
}

// ---- internal ----

func (a *Config) init() error {
	return nil
}

package base

// desc: 基础类

import (
	"fmt"

	"go.uber.org/zap"
)

// Base 所有类都必须继承
type Base struct {
	name   string
	logger *zap.Logger
}

// New ...
func New(name string, logger *zap.Logger) (*Base, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger is nil")
	}

	if len(name) <= 0 {
		return nil, fmt.Errorf("name is empty ")
	}

	return &Base{name: name, logger: logger.Named(name)}, nil
}

// Name ...
func (a *Base) Name() string { return a.name }

// Logger ...
func (a *Base) Logger() *zap.Logger { return a.logger }

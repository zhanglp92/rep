package main

import (
	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/web"
	"go.uber.org/zap"
)

func main() {
	config, err := config.New("rep.yml")
	if err != nil {
		panic(err)
	}

	w, err := web.New(config)
	if err != nil {
		panic(err)
	}

	go w.Run()

	config.Logger().Error("end", zap.Error(<-w.ListenError()))
}

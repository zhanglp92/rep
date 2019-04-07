package main

import (
	_ "github.com/icattlecoder/godaemon"

	"github.com/zhanglp92/rep/config"
	"github.com/zhanglp92/rep/db"
	"github.com/zhanglp92/rep/web"
	"go.uber.org/zap"
)

func main() {
	config, err := config.New("rep.yml")
	if err != nil {
		panic(err)
	}

	ldb, err := db.New(config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := ldb.Close(); err != nil {
			panic(err)
		}
	}()
	db.SetDef(ldb.Op())

	w, err := web.New(config)
	if err != nil {
		panic(err)
	}

	go w.Run()

	config.Logger().Error("end", zap.Error(<-w.ListenError()))
}

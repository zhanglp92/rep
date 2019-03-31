package web

import (
	"fmt"
	"log"
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestWeb(t *testing.T) {

	Index := func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	}

	Hello := func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	}

	// 设置静态目录

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	router.ServeFiles("/lib/*filepath", http.Dir("/Users/clairey/go/src/github.com/zhanglp92/rep/lib"))

	log.Fatal(http.ListenAndServe(":7001", router))
}

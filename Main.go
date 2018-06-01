package main

import (
	"github.com/emicklei/go-restful"
	"QPongServer/http"
	httpG "net/http"
)

func main() {
	restful.Add(http.NewTestingModule())
	err := httpG.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}

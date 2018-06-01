package main

import (
	"github.com/emicklei/go-restful"
	"QPongServer/http"
)

func main() {
	serverInstance := http.NewQPongServer()

	// adding modules
	err := serverInstance.AddModules([]*restful.WebService{
		http.NewTestingModule(),
		http.NewTemplateModule(),
	})
	if err != nil {
		panic(err)
	}

	err = serverInstance.StartServer(serverInstance.ServerConfig)
	if err != nil {
		panic(err)
	}
}

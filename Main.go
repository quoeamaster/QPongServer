package main

import (
	"github.com/emicklei/go-restful"
	"QPongServer/http"
)

// TODO: make it singelton???
var serverInstance http.QPongServerInstance

func main() {
	serverInstance = http.GetQPongServer()

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

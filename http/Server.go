package http

import (
	"net/http"
	"github.com/emicklei/go-restful"
	"fmt"
	"QPongServer/util"
	"sync"
	"context"
	"time"
)


// singleton... MUST be handled here
var syncLock sync.Once
var serverInstance QPongServerInstance

type QPongServerInstance struct {
	ServerConfig *util.Config
	MRequestContext *ModuleRequestContext
}

/**
 *  struct to declare the context for webservice modules to use
 */
type ModuleRequestContext struct {
	Default context.Context
	DefaultCancelFunc context.CancelFunc
	Background context.Context
}

/**
 *  sort of a singleton method to return the only instance of QPongServer
 */
func GetQPongServer() QPongServerInstance {
	syncLock.Do(func() {
		serverInstance = newQPongServer()
	})
	return serverInstance
}

/**
 *  make the init method PRIVATE
 */
func newQPongServer() QPongServerInstance {
	instance := QPongServerInstance{}

	filePtr, err := util.GetConfigFile()
	if err != nil {
		panic(err)
	}
	cfgPtr, err := util.LoadConfigFromFilepath(filePtr.Name())
	if err != nil {
		panic(err)
	}
	instance.ServerConfig = cfgPtr

	// setup the context(s)
	duration60s, err := time.ParseDuration("60s")
	if err != nil {
		panic(err)
	}
	mrc := ModuleRequestContext{}
	mrc.Default, mrc.DefaultCancelFunc = context.WithTimeout(context.Background(), duration60s)
	mrc.Background = context.Background()
	instance.MRequestContext = &mrc

	return instance
}

func (server *QPongServerInstance) AddModules(modules []*restful.WebService) (err error) {
	for _, ws := range modules {
		restful.Add(ws)
	}   // end -- for (modules)
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("something is wrong on adding Modules => %v", r)
		}
	}()
	return err
}

func (server *QPongServerInstance) StartServer(config *util.Config) error {
	serverPortString := fmt.Sprintf(":%v", config.ServerPort)
	fmt.Printf("** QPong server started at %v port **\n", config.ServerPort)

	return http.ListenAndServe(serverPortString, nil)
}

// TODO: add lifecycle hooks like "system halt" "interrupt" etc and call the corresponding service's Cleanup method (e.g. ESConnector.Cleanup)


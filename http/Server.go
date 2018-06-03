package http

import (
	"net/http"
	"github.com/emicklei/go-restful"
	"fmt"
	"QPongServer/util"
	"sync"
	"context"
	"time"
	"os"
	"os/signal"
	"syscall"
	"QPongServer/datastore"
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
	mrc := ModuleRequestContext{}
	mrc.Background = context.Background()
	instance.MRequestContext = &mrc


	// setup signal intercepts
	go serverExitSequence()

	return instance
}

/**
 *  exit sequences for QPon Server shutdown
 *  add lifecycle hooks like "system halt" "interrupt" etc and
 *  call the corresponding service's Cleanup method (e.g. ESConnector.Cleanup)
 */
func serverExitSequence() {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	sig := <- signalChannel
	fmt.Println("sig received =>", sig)
	// call cleanup method(s)
	_, err := datastore.ESPool.ESConnectionPoolCleanup()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(1)
}


/**
 *  helper method to add more WebService module(s)
 */
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

/**
 *  method to start Qpon server
 */
func (server *QPongServerInstance) StartServer(config *util.Config) error {
	serverPortString := fmt.Sprintf(":%v", config.ServerPort)
	fmt.Printf("** QPong server started at %v port **\n", config.ServerPort)

	return http.ListenAndServe(serverPortString, nil)
}



/**
 *  get the context with timeout (60s) plus the cancelFunction (you can use it or wait till 60s timeout)
 */
func (o *ModuleRequestContext) GetDefaultContextAndCancelFunc() (ctx context.Context, cancelFx context.CancelFunc) {
	duration60s, err := time.ParseDuration("60s")
	if err != nil {
		panic(err)
	}
	return context.WithTimeout(context.Background(), duration60s)
}




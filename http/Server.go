/*
 *  Copyright Project - CFactor, Author - quoeamaster, (C) 2018
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package http

import (
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
    "net/http"
)


// singleton... MUST be handled here
var syncLock sync.Once
var serverInstance QPongServerInstance
var wsContainer *restful.Container

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

	// setup the restful.Container
	wsContainer = restful.NewContainer()


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
        //restful.Add(ws)
	    wsContainer.Add(ws)
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

	// setup the cors for this server
    // Add container filter to enable CORS
    cors := restful.CrossOriginResourceSharing {
        ExposeHeaders:  []string{"X-My-Header"},
        AllowedHeaders: []string{"Content-Type", "Accept"},
        AllowedMethods: []string{"GET", "POST"},
        CookiesAllowed: false,
        Container:      wsContainer,
    }
    // Add container filter to respond to OPTIONS
    //wsContainer.Filter(wsContainer.OPTIONSFilter) // not applicable for my case...
    wsContainer.Filter(optionFilterFx)
    wsContainer.Filter(cors.Filter)
    wsContainer.Filter(OriginCheckFilter)

    // TODO: add a new Handler for testing-upload/{productId}... see if works or not...
    //wsContainer.Handle("/testing-upload/{productId}", nil)

    // start server with wsContainer as Handler
    return http.ListenAndServe(serverPortString, wsContainer)
}

/**
 *  self implemented OPTIONS func
 */
func optionFilterFx(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
    if "OPTIONS" != req.Request.Method {
        chain.ProcessFilter(req, resp)
        return
    }

    //archs := req.Request.Header.Get(restful.HEADER_AccessControlAllowHeaders)
    archs := req.Request.Header.Get(restful.HEADER_AccessControlRequestHeaders)
    methods := "POST,GET,DELETE,PUT,OPTIONS,HEAD"
    origin := req.Request.Header.Get(restful.HEADER_Origin)

    resp.AddHeader(restful.HEADER_Allow, methods)
    resp.AddHeader(restful.HEADER_AccessControlAllowHeaders, archs)
    resp.AddHeader(restful.HEADER_AccessControlAllowMethods, methods)
    resp.AddHeader(restful.HEADER_AccessControlAllowOrigin, origin)

    // PS. add logic on OriginCheck here as well??? since should be logical to do so...
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




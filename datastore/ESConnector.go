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

package datastore

import (
	"QPongServer/util"
	"fmt"
	"github.com/olivere/elastic"
)

type ESConnectionPool struct {
	PoolPtr map[string]*ESConnection
}

type ESConnection struct {
	// the client connecting to a certain ES Node
	ClientPtr *elastic.Client
}

/**
 *  connection pool for ElasticSearch
 */
var ESPool = NewESConnectionPool()

/**
 *  create a new connection pool containing ESConnection objects
 */
func NewESConnectionPool() ESConnectionPool {
	pool := ESConnectionPool{}
	pool.PoolPtr = make(map[string]*ESConnection, 1)
	return pool
}

/**
 *  get back an ESConnection instance based on the ES.host value provided
 */
func GetESConnectionByConfig(config *util.Config) (conn *ESConnection, err error) {
	// check if the connection-pool contains the connection
	connPtr := ESPool.PoolPtr[config.ESHost]

	if connPtr != nil {
		conn = connPtr

	} else {
		var clientPtr *elastic.Client

		clientPtr, err = connectToES(config)
		if err != nil {
			return nil, err
		}
		// create a new ESConnection instance and put it back to the pool
		conn = &ESConnection{}
		conn.ClientPtr = clientPtr
		ESPool.PoolPtr[config.ESHost] = conn
	}
	return conn, err
}

func connectToES(config *util.Config) (*elastic.Client, error) {
	clientPtr, err := elastic.NewClient(elastic.SetURL(config.ESHost))
	if err != nil {
		err = fmt.Errorf("something wrong when connecting to ES by host [%v] => %v", config.ESHost, err)
		return nil, err
	}
	return clientPtr, nil
}

/**
 *  cleanup method (called when the server is going to die)
 */
func (o *ESConnectionPool) ESConnectionPoolCleanup() (ok bool, err error) {
	if o.PoolPtr != nil && len(o.PoolPtr)>0 {
		for _, esConn := range o.PoolPtr {
			ok, err = IsESConnValid(esConn)
			if ok == true {
				esConn.ClientPtr = nil
			}   // end -- esConn and esConn.ClientPtr != nil
		}   // end -- for (iterate the pool's esConn)
	}   // end -- if (poolPtr is valid)
	return ok, err
}

/**
 *  check is ESConnection instance is valid or not
 */
func IsESConnValid(esConn *ESConnection) (bool, error) {
	if esConn != nil && esConn.ClientPtr != nil {
		return true, nil
	} else {
		return false, fmt.Errorf("esConn is INVALID [%v]", esConn.ClientPtr)
	}
}

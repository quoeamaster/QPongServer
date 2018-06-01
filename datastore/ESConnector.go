package datastore

import (
	"QPongServer/util"
	"github.com/elastic/go-elasticsearch/client"
	"fmt"
)

type ESConnectionPool struct {
	PoolPtr map[string]*ESConnection
}

type ESConnection struct {
	// the client connecting to a certain ES Node
	ClientPtr *client.Client
}

/**
 *  connection pool for ElasticSearch
 */
var ESPool ESConnectionPool = NewESConnectionPool()

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
		clientPtr, err := client.New(client.WithHost(config.ESHost))
		if err != nil {
			err = fmt.Errorf("something wrong when connecting to ES by host [%v] => %v", config.ESHost, err)
		}
		// create a new ESConnection instance and put it back to the pool
		conn = &ESConnection{}
		conn.ClientPtr = clientPtr
		ESPool.PoolPtr[config.ESHost] = conn
	}
	return conn, err
}
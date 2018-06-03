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
func (o *ESConnectionPool) Cleanup() (ok bool, err error) {
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

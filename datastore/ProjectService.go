package datastore

import (
	"github.com/elastic/go-elasticsearch/api"
	"fmt"
)

const IndexProject = "qpong_project"
const DocType =      "doc"

// TODO: a mechanism to create settings and mappings for the 1st time as well...

func PersistProjectEntity(p *Project, esConn *ESConnection) (resp *api.IndexResponse, err error) {
	// use overwrite mechanism (not update api from es; just INDEX)
	if esConn != nil && esConn.ClientPtr != nil && p != nil {
		client := *(esConn.ClientPtr)
		body := map[string]interface{} {
			"project": *p,
		}
		resp, err = client.Index(IndexProject, DocType,body)
		if err != nil {
			return resp, err
		}
		// check for logical problems
fmt.Println("**", *resp, (*resp).Response.Status, (*resp).Response.StatusCode)

	}   // end -- if (esConn, clientPtr and p is valid)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return resp, err
}

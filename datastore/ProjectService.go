package datastore

import (
	"fmt"
	"context"
	"github.com/olivere/elastic"
)

const IndexProject = "qpong_project"
const DocType =      "doc"

// TODO: a mechanism to create settings and mappings for the 1st time as well...

func PersistProjectEntity(p *Project, esConn *ESConnection) (iResp *elastic.IndexResponse, err error) {
	// use overwrite mechanism (not update api from es; just INDEX)
	if esConn != nil && esConn.ClientPtr != nil && p != nil {
		client := esConn.ClientPtr
		iResp, err = client.Index().Index("qpon_project").
			Type("doc").Pretty(true).BodyJson(*p).
			Do(context.Background())
		if err != nil {
			return iResp, err
		}
		// logical checks (such as response code is 500)
		fmt.Println("response =>", iResp.Result)

	}   // end -- if (esConn, clientPtr and p is valid)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	return iResp, err
}

// to cat indices...
/*
catISrv := client.CatIndices().Pretty(true).Columns("index", "health").Sort("index", "health")
cResp, err := catISrv.Do(context.Background())
if err != nil {
return resp, err
}
// check for logical problems
cRespRows := []elastic.CatIndicesResponseRow(cResp)
fmt.Println(len(cRespRows))
fmt.Println(cRespRows[0].Index, cRespRows[0].Health)
fmt.Println(cRespRows[1].Index, cRespRows[1].Health)
fmt.Println(cRespRows[10].Index, cRespRows[10].Health)
*/
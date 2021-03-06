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
	"fmt"
	"github.com/olivere/elastic"
	"context"
)

const IndexProject = "qpong_project"
const DocType =      "doc"

func createIndexProjectSrv(esConn *ESConnection) *elastic.IndexService {
	return esConn.ClientPtr.Index().
		Index(IndexProject).Type(DocType).Pretty(true)
}

// TODO: a mechanism to create settings and mappings for the 1st time as well...

func PersistProjectEntity(p *Project,
	esConn *ESConnection,
	timeoutCtx context.Context,
	timeoutCtxCancelFx context.CancelFunc) (iResp *elastic.IndexResponse, err error) {

	// use overwrite mechanism (not update api from es; just INDEX)
	var valid bool

	valid, err = IsESConnValid(esConn)
	if valid {
		if p != nil {
			iResp, err = createIndexProjectSrv(esConn).BodyJson(*p).
				Do(timeoutCtx)
			if err != nil {
				return iResp, err
			}
			// logical checks (such as response code is 500)
			fmt.Println("response =>", iResp.Result)
		} else {
			err = fmt.Errorf("project entity provied is nil~ [%v]", *p)
		}
	}
	/* else {
		err = fmt.Errorf("esConn is INVALID [%v]", esConn.ClientPtr)
	}*/   // end -- if (esConn, clientPtr and p is valid)

	defer func() {
		if timeoutCtxCancelFx != nil {
			timeoutCtxCancelFx()
		}
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
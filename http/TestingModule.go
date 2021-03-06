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
	"QPongServer/util"
	"fmt"
)

/**
 *  model for testing
 */
type TestingModel struct {
	ArtistId string
	ResourceName string     // resource name / identifier (unique within the Artist)
	ResourceBytes []byte    // resource (binary) in []byte
}

func NewTestingModule() *restful.WebService {
	srv := new(restful.WebService)
	srv.Path("/testing").
		Consumes(restful.MIME_JSON, restful.MIME_XML, "multipart/form-data").
		Produces(restful.MIME_JSON)

	// test on GET, POST route
	srv.Route(srv.GET("/{productId}").To(getProductById)).
		Route(srv.POST("/{productId}").To(postProductIdUpload))

	return srv
}

/*
func debugOriginFromHeader(header http.Header) {
	origin, err := util.GetOriginFromHeaders(header)
	if err != nil {
		panic(err)
	}
	fmt.Println("protocol=>",origin.Protocol,"host=>",origin.Host, "port=>", origin.Port, "fullAddr:", origin.Address)
} */

func getProductById(req *restful.Request, res *restful.Response) {
	pId := req.PathParameter("productId")
	testingModel := TestingModel{}

	//debugOriginFromHeader(req.Request.Header)

	switch pId {
	case "101":
		testingModel.ArtistId = "101"
		testingModel.ResourceName = "painting_101"
	default:
		testingModel.ArtistId = "unknown"
		testingModel.ResourceName = "unknown"
	}
    err := res.WriteHeaderAndJson(200, testingModel, restful.MIME_JSON)
    if err != nil {
        panic(err)
    }
}

func postProductIdUpload(req *restful.Request, res *restful.Response) {
    fmt.Println("** inside postProductIdUpload")

    pfile, fileHeaderPtr, err := req.Request.FormFile("pfile")
    if err != nil {
        res.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
        return
    }
    if fileHeaderPtr != nil {
        fmt.Println(fileHeaderPtr.Filename, fileHeaderPtr.Size, fileHeaderPtr.Header)
    }
    if pfile != nil {
        iBytes, err := util.WriteMultiPartFileToDisc(&pfile,
            util.NewStorageMeta(GetQPongServer().ServerConfig.ServerDataPath+"/"+fileHeaderPtr.Filename))

        if err != nil {
            fmt.Println("bb write to disc failed")
            res.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
            return
        } else {
            res.WriteHeaderAndJson(200, fmt.Sprintf("successfully writen the file %v with bytes: %v", fileHeaderPtr.Filename, iBytes), restful.MIME_JSON)
            return
        }
    }

    /* if using FormFile, don't use parseMultipartForm(mem)
    err := req.Request.ParseMultipartForm(1000000*5)
    if err != nil {
        fmt.Print("parse multipart form failed")
        res.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
        return
    }
    fmt.Println("multipartForm", req.Request.MultipartForm)
    */

    /* if used parseMultipartForm(mem); don't try to get MultipartReader...
    readerPtr, err := req.Request.MultipartReader()
    if err != nil {
        fmt.Print("get MultipartReader failed")
        res.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
        return
    }
    formPtr, err := readerPtr.ReadForm(128)
    if err != nil {
        fmt.Print("read form failed")
        res.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
        return
    }
    fmt.Println(formPtr.File, " vs", formPtr.Value)
    */

    res.WriteHeaderAndJson(200, "everything seems ok", restful.MIME_JSON)
}




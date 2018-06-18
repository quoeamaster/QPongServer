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
	"QPongServer/datastore"
    "fmt"
    "QPongServer/util"
)

// struct for encapsulating the request parameters
type TemplateDataModel struct {
    ProjectId string
    Title string
    Subtitle string
    Description string
    PickedImageList []string
    PickedCategoryList []string
}

// create an instance of TemplateDataModel struct based on the provided
// data provided
func NewTemplateDataModel(projectId, title, subtitle, description, pickedImageListString, pickedCategoryListString string) TemplateDataModel  {
    m := TemplateDataModel{}

    if !util.IsStringEmpty(projectId) {
        m.ProjectId = projectId
    }
    if !util.IsStringEmpty(title) {
        m.Title = title
    }
    if !util.IsStringEmpty(subtitle) {
        m.Subtitle = subtitle
    }
    if !util.IsStringEmpty(description) {
        m.Description = description
    }
    // TODO: convert the string back to array

    return m
}

/**
 *  creation of the TemplateModule
 */
func NewTemplateModule() *restful.WebService {
	srv := new(restful.WebService)
	srv.Path("/template").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	srv.Route(srv.POST("/generate/{project-id}").To(generateTemplateForProject)).
		Route(srv.GET("/suggestLayout").To(suggestLayoutWithProjectId))

	return srv
}

func generateTemplateForProject(req *restful.Request, res *restful.Response) {
	projectId := req.PathParameter("project-id")
	projectInstancePtr := datastore.NewProjectEntity(projectId)

	// add design etc and test connectivity to ES (based on a closed ES, see what would happen)
	projectInstancePtr.AddDesign(
		datastore.NewDesignEntity().AddSpec(
			datastore.NewSpecEntity().
				AddBackgroundImagePath("abc.jpg").
				AddDescription(datastore.NewTextBlockEntity().AddText("hi")).
				AddSubTitle(datastore.NewTextBlockEntity().AddText("world"))))
	// get esConnection
	esConn, err := datastore.GetESConnectionByConfig(GetQPongServer().ServerConfig)
	if err != nil {
		res.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
		return
	}
	// get the default context + cancel fx
	ctx, ctxCancel := GetQPongServer().MRequestContext.GetDefaultContextAndCancelFunc()
	iResp, err := datastore.PersistProjectEntity(projectInstancePtr,
		esConn,
		ctx, ctxCancel)
		// ** OR the background context... GetQPongServer().MRequestContext.Background, nil)
	if err != nil {
		res.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
		return
	}
	res.WriteHeaderAndJson(200, iResp, restful.MIME_JSON)
}

// method to get suggestion on LAYOUT based on the given data:
// 1) title, subtitle and description;
// 2) the picked image(s)
// 3) theoretically, the suggestions should come from some AI or ML results
//      (but due to the fact that the project is still young; at the meantime
//      only simple if then else logic would pretend to the smart brain)
func suggestLayoutWithProjectId(req *restful.Request, res *restful.Response)  {
    // this projectId should be saved later on...
    // TODO: (caching by "projectId" and "suggestionId" etc)
    projectId := req.QueryParameter("projectId")
    fmt.Println(req.QueryParameter("pickedImages"))
    fmt.Println(req.QueryParameter("pickedCategories"))

    res.WriteHeaderAndJson(200,
        NewModuleResponse( fmt.Sprintf("testing only %v", projectId)),
        restful.MIME_JSON)
}

package http

import (
	"github.com/emicklei/go-restful"
	"QPongServer/datastore"
	"fmt"
)

/**
 *  creation of the TemplateModule
 */
func NewTemplateModule() *restful.WebService {
	srv := new(restful.WebService)
	srv.Path("/template").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	srv.Route(srv.POST("/generate/{project-id}").To(generateTemplateForProject))

	return srv
}

func generateTemplateForProject(req *restful.Request, res *restful.Response) {
	projectId := req.PathParameter("project-id")
	projectInstancePtr := datastore.NewProjectEntity(projectId)
	// TODO: add design etc and test connectivity to ES (based on a closed ES, see what would happen)
	projectInstancePtr.AddDesign(
		datastore.NewDesignEntity().AddSpec(
			datastore.NewSpecEntity().
				AddBackgroundImagePath("abc.jpg").
				AddDescription(datastore.NewTextBlockEntity().AddText("hi")).
				AddSubTitle(datastore.NewTextBlockEntity().AddText("world"))))
	// get esConnection
	esConn, err := datastore.GetESConnectionByConfig(GetQPongServer().ServerConfig)
	if err != nil {
		res.WriteHeaderAndJson(500, PopulateModuleError(&err), restful.MIME_JSON)
		return
	}
	iResp, err := datastore.PersistProjectEntity(projectInstancePtr, esConn)
	if err != nil {
		res.WriteHeaderAndJson(500, PopulateModuleError(&err), restful.MIME_JSON)
		return
	}
	fmt.Println("## responses =>", iResp)

	/*
	response, err := esConn.ClientPtr.Cat.Indices()
	fmt.Println("** can do CAT?")
	if err != nil {
		meta := make(map[string]interface{}, 1)
		origin, err2 := util.GetOriginFromHeaders(req.Request.Header)
		if err2 != nil {
			panic(err2)
		}
		meta["ip"] = origin.String()

		errM := PopulateModuleError(&err, meta)
		res.WriteHeaderAndJson(500, errM, restful.MIME_JSON)
		return
	}
	var bBytes []byte
	response.Response.Body.Read(bBytes)
	fmt.Println(string(bBytes))
	*/
}

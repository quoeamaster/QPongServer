package http

import (
	"github.com/emicklei/go-restful"
	"QPongServer/datastore"
)

/**
 *  creation of the TemplateModule
 */
func NewTemplateModule() *restful.WebService {
	srv := new(restful.WebService)
	srv.Path("/template").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	// add security filter(s) to the service
	srv.Filter(OriginCheckFilter)

	srv.Route(srv.POST("/generate/{project-id}").To(generateTemplateForProject))

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
		res.WriteHeaderAndJson(500, PopulateModuleError(&err), restful.MIME_JSON)
		return
	}
	// get the default context + cancel fx
	ctx, ctxCancel := GetQPongServer().MRequestContext.GetDefaultContextAndCancelFunc()
	iResp, err := datastore.PersistProjectEntity(projectInstancePtr,
		esConn,
		ctx, ctxCancel)
		// ** OR the background context... GetQPongServer().MRequestContext.Background, nil)
	if err != nil {
		res.WriteHeaderAndJson(500, PopulateModuleError(&err), restful.MIME_JSON)
		return
	}
	res.WriteHeaderAndJson(200, iResp, restful.MIME_JSON)
}

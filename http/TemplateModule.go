package http

import (
	"github.com/emicklei/go-restful"
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
	fmt.Println("project-id =>", req.PathParameter("project-id"))
}

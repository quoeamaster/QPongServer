package http

import (
	"github.com/emicklei/go-restful"
	"net/http"
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
		Consumes(restful.MIME_JSON, restful.MIME_XML, restful.MIME_OCTET).
		Produces(restful.MIME_JSON)

	// test on GET, POST route
	srv.Route(srv.GET("/{productId}").To(getProductById)).
		Route(srv.POST("").To(postProductsUpdateToAll))

	return srv
}

func debugOriginFromHeader(header http.Header) {
	origin, err := util.GetOriginFromHeaders(header)
	if err != nil {
		panic(err)
	}
	fmt.Println("protocol=>",origin.Protocol,"host=>",origin.Host, "port=>", origin.Port, "fullAddr:", origin.Address)
}

func getProductById(req *restful.Request, res *restful.Response) {
	pId := req.PathParameter("productId")
	testingModel := TestingModel{}

	debugOriginFromHeader(req.Request.Header)

	switch pId {
	case "101":
		testingModel.ArtistId = "101"
		testingModel.ResourceName = "painting_101"
	default:
		testingModel.ArtistId = "unknown"
		testingModel.ResourceName = "unknown"
	}
	res.AddHeader("Access-Control-Allow-Origin", "http://localhost:8080")
	err := res.WriteHeaderAndJson(200, testingModel, restful.MIME_JSON)
	if err != nil {
		panic(err)
	}
}

func postProductsUpdateToAll(req *restful.Request, res *restful.Response) {

}




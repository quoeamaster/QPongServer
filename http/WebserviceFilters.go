package http

import (
    "strings"
    "fmt"
    "QPongServer/util"
    "github.com/emicklei/go-restful"
)

const HeaderAccessControlAllowOrigin = "Access-Control-Allow-Origin"

/**
 *  filter to check if the request Origin is allowed for accessing the server's api / endpoint
 */
func OriginCheckFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
    origin, err := util.GetOriginFromHeaders(req.Request.Header)
    if err != nil {
        resp.WriteHeaderAndJson(500, PopulateModuleError(&err), restful.MIME_JSON)
        return
    }
    _, err = isOriginAllowed(origin, GetQPongServer().ServerConfig.AllowedAccessList)
    if err != nil {
        resp.WriteHeaderAndJson(500, PopulateModuleError(&err), restful.MIME_JSON)
        return

    } else {
        // add back the corresponding Access-Control-Allow-Origin header
        resp.AddHeader(HeaderAccessControlAllowOrigin, origin.Address)

        // everything is fine, forward to the next "filter"
        chain.ProcessFilter(req, resp)
    }
}

/**
 *  helper method to check if the provided origin could
 *  access the server's features.
 *  the allowed origin address list is configured in config files or other sources
 */
func isOriginAllowed(origin util.Origin, allowedList []string) (bool, error) {
    valid := false
    var err error = nil

    if allowedList != nil {
        for _, site := range allowedList {
            if strings.Compare(site, origin.Address) == 0 {
                valid = true
                break
            }   // end -- if (origin.Address == site)
        }   // end -- for (allowedList)
    }
    // if NOT allowed... create an error
    if !valid {
        err = fmt.Errorf("[%v] is not a member on the allowed access list. Please check the %v", origin.Address, util.QpongDefaultConfFilepath)
    }
    return valid, err
}
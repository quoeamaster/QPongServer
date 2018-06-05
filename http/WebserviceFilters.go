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
    "strings"
    "fmt"
    "QPongServer/util"
    "github.com/emicklei/go-restful"
)

//const HeaderAccessControlAllowOrigin = "Access-Control-Allow-Origin"

/**
 *  filter to check if the request Origin is allowed for accessing the server's api / endpoint
 */
func OriginCheckFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
    //fmt.Println("** inside security fileter")
    origin, err := util.GetOriginFromHeaders(req.Request.Header)
    if err != nil {
        resp.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
        return
    }
    _, err = isOriginAllowed(origin, GetQPongServer().ServerConfig.AllowedAccessList)
    if err != nil {
        resp.WriteHeaderAndJson(500, NewModuleError(&err), restful.MIME_JSON)
        return

    } else {
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
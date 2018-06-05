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

package util

import (
	"net/http"
	"strings"
	"strconv"
	"fmt"
)

type Origin struct {
	Address string
	Host string
	Port int
	Protocol string // http or https etc
}

func (o *Origin) String() string {
	return fmt.Sprintf("address: %v", o.Address)
}

/**
 *  method to extract the Origin data if available
 */
func GetOriginFromHeaders(header http.Header) (origin Origin, err error) {
	origin = Origin{}

	if header != nil {
		addr := header.Get("Origin")
		if strings.Compare("", strings.TrimSpace(addr)) != 0 {
			origin.Address = addr

			parts := strings.Split(addr, ":")
			// e.g. http : //localhost : 8080
			if len(parts) == 3 {
				origin.Port, err = strconv.Atoi(parts[2])
				origin.Protocol = parts[0]
				// removing "//"
				origin.Host = parts[1][2:]
			}   // end -- if (http://localhost , 8080)
		}   // end -- if (non empty string)
	}
	return origin, err
}


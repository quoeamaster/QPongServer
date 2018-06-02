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

/**
 * TODO:
 *  helper method to check if the provided origin could
 *  access the server's features.
 *  the allowed origin address list is configured in config files or other sources
 */
func IsOriginAllowed(origin Origin) (bool, error) {

	return false, nil
}
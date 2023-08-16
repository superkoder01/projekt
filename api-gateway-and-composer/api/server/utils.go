/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package server

import (
	"github.com/op/go-logging"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/api-gateway-and-composer/internal/error"
	"strings"
)

const (
	space  = " "
	bearer = "Bearer"
)

var (
	logger = logging.MustGetLogger("http_server")
)

func getTokenFromAuthHeader(header string) (string, error) {
	header = strings.TrimSpace(header)
	if header == "" {
		return "", e.ApiErrNoAuthorizationHeader
	}

	headerContent := strings.Split(header, space)
	if len(headerContent) != 2 || headerContent[0] != bearer {
		return "", e.ApiErrAuthorizationHeaderInvalid
	}

	return headerContent[1], nil
}

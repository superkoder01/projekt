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

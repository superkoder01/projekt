package proxy

import (
	"strings"
)

func trimContextUrlPathSuffix(s string) string {
	return strings.TrimSuffix(s, "/")
}
func getProxyAddress(s string) string {
	paths := strings.Split(s, "/")
	return "/" + paths[1]
}

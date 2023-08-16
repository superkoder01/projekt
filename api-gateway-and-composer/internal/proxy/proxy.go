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
package proxy

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strings"
)

type proxy struct {
	url *url.URL
}

const ()

var (
	logger  = logging.MustGetLogger("proxy")
	proxies map[string]*proxy
)

func RegisterProxies(r *gin.Engine) error {
	proxies = make(map[string]*proxy)
	//viper.SetConfigFile(*configPath)

	//if err := viper.ReadInConfig(); err != nil {
	//	return err
	//}

	setupProxyHandlers()
	runProxies(r)
	return nil
}

func setupProxyHandlers() {
	servers := viper.GetStringMapString("server")
	for serviceName, _ := range servers {
		service := fmt.Sprintf("server.%s", serviceName)
		address := viper.GetString(service + ".address")
		url, err := url.Parse(address)
		if err != nil {
			logger.Errorf("Error parsing url %s with error %v", address, err)
			continue
		}
		for _, path := range viper.GetStringSlice(service + ".path") {
			url.Path = path
			proxies[path] = &proxy{url: url}
		}
	}
}

func runProxies(r *gin.Engine) {
	for path := range proxies {
		r.Any(path, runProxy)
	}
}

func runProxy(context *gin.Context) {
	reqPath := context.Request.URL.Path
	foundKey := ""
	for k := range proxies {
		r := strings.Replace(k, "*any", ".*", -1)
		if matched, err := regexp.MatchString(r, reqPath); err == nil {
			if matched {
				foundKey = k
				break
			}
		} else {
			logger.Errorf("Error parsing proxy pattern %s with error %v", k, err)
		}
	}

	if proxy, ok := proxies[foundKey]; ok {
		server := httputil.NewSingleHostReverseProxy(proxy.url)
		server.Director = func(request *http.Request) {
			request.Header = context.Request.Header
			request.Host = proxy.url.Host
			request.URL.Host = proxy.url.Host
			request.URL.Scheme = proxy.url.Scheme
			request.URL.Path = context.Request.URL.Path
		}
		server.ServeHTTP(context.Writer, context.Request)
	} else {

	}

}

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
package http_utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"net"
	"net/http"
)

const (
	HTTP = "http://"
	GET  = "GET"
	POST = "POST"
)

func Get(ctx *gin.Context, host string, port string, url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(GET, HTTP+net.JoinHostPort(host, port)+url, nil)
	if err != nil {
		e.HandleError(err, ctx)
	}
	req.Header = ctx.Request.Header
	resp, _ := client.Do(req)
	return resp
}

func Post(ctx *gin.Context, host string, port string, url string, body model.Model) *http.Response {
	client := &http.Client{}
	reqBody, err := json.Marshal(body)
	req, err := http.NewRequest(POST, HTTP+net.JoinHostPort(host, port)+url, bytes.NewBuffer(reqBody))
	if err != nil {
		e.HandleError(err, ctx)
	}
	req.Header = ctx.Request.Header
	resp, _ := client.Do(req)
	return resp
}

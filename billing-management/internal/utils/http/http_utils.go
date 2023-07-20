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

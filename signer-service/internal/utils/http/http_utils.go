package http_utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/error"
	"net"
	"net/http"
)

const (
	HTTP  = "http://"
	GET   = "GET"
	POST  = "POST"
	PUT   = "PUT"
	PATCH = "PATCH"
)

func Get(ctx *gin.Context, host string, port string, url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(GET, HTTP+net.JoinHostPort(host, port)+url, nil)
	if err != nil {
		e.HandleError(err, ctx)
	}
	req.Header = ctx.Request.Header
	resp, err := client.Do(req)
	if err != nil {
		e.HandleError(err, ctx)
	}
	return resp
}

func Patch(ctx *gin.Context, host string, port string, url string) *http.Response {
	client := &http.Client{}
	req, err := http.NewRequest(PATCH, HTTP+net.JoinHostPort(host, port)+url, nil)
	if err != nil {
		e.HandleError(err, ctx)
	}
	req.Header = ctx.Request.Header
	resp, _ := client.Do(req)
	return resp
}

func Put(ctx *gin.Context, host string, port string, url string, body interface{}) *http.Response {
	client := &http.Client{}
	reqBody, err := json.Marshal(body)
	req, err := http.NewRequest(PUT, HTTP+net.JoinHostPort(host, port)+url, bytes.NewBuffer(reqBody))
	if err != nil {
		e.HandleError(err, ctx)
	}
	req.Header = ctx.Request.Header
	resp, _ := client.Do(req)
	return resp
}

func Post(ctx *gin.Context, host string, port string, url string, body interface{}) *http.Response {
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

package mock

import "net/http"

type MockHttpClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (receiver MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	return receiver.DoFunc(req)
}

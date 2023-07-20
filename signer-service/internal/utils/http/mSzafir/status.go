package mSzafir

import "fmt"

type StatusResponse struct {
	Completed  bool     `json:"completed"`
	StatusCode int32    `json:"statusCode"`
	Status     string   `json:"status"`
	FileUrls   []string `json:"fileUrls"`
}

func (w *StatusResponse) String() string {
	return fmt.Sprintf("%s", *w)
}

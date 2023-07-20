package mSzafir

import "fmt"

type InitResponse struct {
	SigningId  string   `json:"signingId"`
	StatusCode int32    `json:"statusCode"`
	Status     string   `json:"status"`
	FileIds    []string `json:"fileIds"`
}

func (w *InitResponse) String() string {
	return fmt.Sprintf("%s", *w)
}

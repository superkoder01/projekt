package mSzafir

import "fmt"

type LoginHauth struct {
	Timestamp int64    `json:"timestamp"`
	SigningId string   `json:"signingId"`
	FileIds   []string `json:"fileIds"`
	Hmac      string   `json:"hmac"`
}

func (w *LoginHauth) String() string {
	return fmt.Sprintf("%s", *w)
}

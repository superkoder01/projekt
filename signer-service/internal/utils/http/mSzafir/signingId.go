package mSzafir

import "fmt"

type SigningId struct {
	SigningId string `json:"signingId"`
}

func (w *SigningId) String() string {
	return fmt.Sprintf("%s", *w)
}

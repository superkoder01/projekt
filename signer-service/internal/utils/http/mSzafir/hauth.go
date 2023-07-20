package mSzafir

import "fmt"

type Hauth struct {
	Timestamp                       int64  `json:"timestamp"`
	Mode                            string `json:"mode"`
	UrlSigningCompleted             string `json:"urlSigningCompleted"`
	UrlSigningCompletedNotification string `json:"urlSigningCompletedNotification"`
	Hmac                            string `json:"hmac"`
}

func (w *Hauth) String() string {
	return fmt.Sprintf("%s", *w)
}

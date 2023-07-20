package mSzafir

import "fmt"

type Init struct {
	Files []byte `json:"files"`
}

func (w *Init) String() string {
	return fmt.Sprintf("%s", *w)
}

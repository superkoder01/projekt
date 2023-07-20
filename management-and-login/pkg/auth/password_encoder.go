package auth

import (
	"golang.org/x/crypto/bcrypt"
)

type Encoder interface {
	Encode(string) (string, error)
	Compare(hash string, pass string) error
}

type encoder struct {
}

func newEncoder() *encoder {
	return &encoder{}
}

func (e *encoder) Encode(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (e *encoder) Compare(hash string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
}

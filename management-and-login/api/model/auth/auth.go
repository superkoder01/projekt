package auth

import "fmt"

type Auth struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

func NewAuth() *Auth {
	return &Auth{}
}

func (l *Auth) String() string {
	return fmt.Sprintf("%s", *l)
}

// SETTERS

func (l *Auth) SetLogin(s string) {
	l.Login = s
}

func (l *Auth) SetPassword(s string) {
	l.Password = s
}

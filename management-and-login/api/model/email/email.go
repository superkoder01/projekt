package email

import "fmt"

type Email struct {
	Email string `json:"email,omitempty"`
}

func NewEmail() *Email {
	return &Email{}
}

func (e *Email) String() string {
	return fmt.Sprintf("%s", *e)
}

package administrator

import (
	"fmt"
	"time"
)

type Administrator struct {
	ProviderID int `json:"providerId,omitempty"`

	ID int `json:"id,omitempty"`

	Active bool `json:"isActive,omitempty"`

	Login *string `json:"login"`

	Email string `json:"email,omitempty"`

	RoleID int `json:"roleId,omitempty"`

	WorkerID int `json:"workerId,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Phone string `json:"phone,omitempty"`

	ExtraInfo string `json:"extraInfo,omitempty"`

	AddedDate time.Time `json:"addedDate,omitempty"`
}

func NewAdministrator() *Administrator {
	return &Administrator{}
}

func (a *Administrator) String() string {
	return fmt.Sprintf("%s", *a)
}

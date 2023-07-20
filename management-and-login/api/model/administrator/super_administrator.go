package administrator

import (
	"fmt"
	"time"
)

type SuperAdministrator struct {
	ProviderID int `json:"providerId,omitempty"`

	ID int `json:"id,omitempty"`

	Active bool `json:"isActive,omitempty"`

	Login *string `json:"login"`

	Email string `json:"email,omitempty"`

	RoleID int `json:"roleId,omitempty"`

	WorkerID int `json:"workerId,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`

	AddedDate time.Time `json:"addedDate,omitempty"`
}

func NewSuperAdministrator() *SuperAdministrator {
	return &SuperAdministrator{}
}

func (a *SuperAdministrator) String() string {
	return fmt.Sprintf("%s", *a)
}

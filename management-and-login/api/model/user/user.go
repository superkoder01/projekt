package user

import (
	"fmt"
	"time"
)

type User struct {

	// customer_account Id
	CustomerAccountID int `json:"customerAccountId,omitempty"`

	// company Id
	ProviderID int `json:"providerId,omitempty"`

	// id
	// Read Only: true
	ID int `json:"id,omitempty"`

	// is active
	Active bool `json:"isActive,omitempty"`

	// login
	// Example: user101
	// Required: true
	Login *string `json:"login"`

	Email string `json:"email,omitempty"`

	// must change password
	MustChangePassword bool `json:"mustChangePassword,omitempty"`

	// password
	// Example: pass101!
	// Required: true
	Password *string `json:"password"`

	// role Id
	RoleID int `json:"roleId,omitempty"`

	// worker Id
	WorkerID int `json:"workerId,omitempty"`

	// added date
	AddedDate time.Time `json:"addedDate,omitempty"`
}

type UserOption func(c *User)

func NewUser(opts ...UserOption) *User {
	user := &User{}

	for _, option := range opts {
		option(user)
	}

	return user
}

func (u *User) String() string {
	return fmt.Sprintf("%s", *u)
}

// Constructor options

func Login(s *string) UserOption {
	return func(u *User) {
		u.Login = s
	}
}

func Password(s *string) UserOption {
	return func(u *User) {
		u.Password = s
	}
}

func ProviderId(i int) UserOption {
	return func(u *User) {
		u.ProviderID = i
	}
}

func ClientId(i int) UserOption {
	return func(u *User) {
		u.CustomerAccountID = i
	}
}

func RoleId(i int) UserOption {
	return func(u *User) {
		u.RoleID = i
	}
}

func Active(b bool) UserOption {
	return func(u *User) {
		u.Active = b
	}
}

func MustChangePassword(b bool) UserOption {
	return func(u *User) {
		u.MustChangePassword = b
	}
}

func WorkerId(i int) UserOption {
	return func(u *User) {
		u.WorkerID = i
	}
}

// SETTERS

func (u *User) SetLogin(s *string) {
	u.Login = s
}

func (u *User) SetPassword(s *string) {
	u.Password = s
}

func (u *User) SetProviderId(i int) {
	u.ProviderID = i
}

func (u *User) SetClientId(i int) {
	u.CustomerAccountID = i
}

func (u *User) SetRoleId(i int) {
	u.RoleID = i
}

func (u *User) SetActive(b bool) {
	u.Active = b
}

func (u *User) SetMustChangePassword(b bool) {
	u.MustChangePassword = b
}

func (u *User) SetWorkerId(i int) {
	u.WorkerID = i
}

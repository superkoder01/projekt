package user_change_password

import "fmt"

type UserChangePassword struct {
	OldPassword       string `json:"oldPassword,omitempty"`
	NewPassword       string `json:"newPassword,omitempty"`
	NewPasswordRetype string `json:"newPassword_retype,omitempty"`
}

func NewUserChangePassword() *UserChangePassword {
	return &UserChangePassword{}
}

func (u *UserChangePassword) String() string {
	return fmt.Sprintf("%s", *u)
}

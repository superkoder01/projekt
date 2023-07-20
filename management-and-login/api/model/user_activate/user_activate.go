package user_activate

import "fmt"

type UserActivate struct {
	NewPassword       string `json:"newPassword"`
	NewPasswordRetype string `json:"newPasswordRetype"`
}

func NewUserActivate() *UserActivate {
	return &UserActivate{}
}

func (u *UserActivate) String() string {
	return fmt.Sprintf("%s", *u)
}

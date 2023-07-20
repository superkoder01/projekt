package logout

import "fmt"

type Logout struct {
	AccessID  string `json:"accessId"`
	RefreshID string `json:"refreshId"`
}

func NewLogout() *Logout {
	return &Logout{}
}

func (l *Logout) String() string {
	return fmt.Sprintf("%s", *l)
}

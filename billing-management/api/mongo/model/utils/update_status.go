package utils

import (
	"fmt"
)

type UpdateStatus struct {
	Status string `json:"status"`
}

func (c *UpdateStatus) String() string {
	return fmt.Sprintf("%s", *c)
}

package parameter_name

import (
	"fmt"
)

type ParameterName struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
}

func (c *ParameterName) String() string {
	return fmt.Sprintf("%s", *c)
}

func (c *ParameterName) SetName(s string) {
	c.Name = s
}

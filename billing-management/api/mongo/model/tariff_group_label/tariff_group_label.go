package tariff_group_label

import "fmt"

type TariffGroupLabel struct {
	Id   string `json:"id" bson:"_id"`
	Name string `bson:"name"`
}

func (c *TariffGroupLabel) String() string {
	return fmt.Sprintf("%s", *c)
}

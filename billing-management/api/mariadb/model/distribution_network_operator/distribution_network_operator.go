package distribution_network_operator

import "fmt"

type DistributionNetworkOperator struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Nip     string `json:"nip"`
	Address string `json:"address"`
	City    string `json:"city"`
}

func (c *DistributionNetworkOperator) String() string {
	return fmt.Sprintf("%s", *c)
}

func (c *DistributionNetworkOperator) SetCity(s string) {
	c.City = s
}

func (c *DistributionNetworkOperator) SetAddress(s string) {
	c.Address = s
}

func (c *DistributionNetworkOperator) SetNip(s string) {
	c.Nip = s
}

func (c *DistributionNetworkOperator) SetName(s string) {
	c.Name = s
}

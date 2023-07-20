package utils

type GetByProviderId struct {
	ProviderID int
}

func (c *GetByProviderId) SetProviderID(i int) {
	c.ProviderID = i
}

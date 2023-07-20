package utils

type GetByProviderIdAndAccountId struct {
	AccountID  int
	ProviderID int
}

func (c *GetByProviderIdAndAccountId) SetProviderID(i int) {
	c.ProviderID = i
}

func (c *GetByProviderIdAndAccountId) SetAccountID(i int) {
	c.AccountID = i
}

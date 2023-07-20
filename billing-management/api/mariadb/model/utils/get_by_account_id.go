package utils

type GetByAccountId struct {
	AccountID int
}

func (c *GetByAccountId) SetAccountID(i int) {
	c.AccountID = i
}

package billing

import "time"

type DepositItemValue struct {
	Amount int `json:"amount" bson:"amount"`
	Value  int `json:"value" bson:"value"`
}

type DepositItem struct {
	In       DepositItemValue `json:"in" bson:"in"`
	Out      DepositItemValue `json:"out" bson:"out"`
	Has      DepositItemValue `json:"has" bson:"has"`
	Used     DepositItemValue `json:"used" bson:"used"`
	Residual DepositItemValue `json:"residual" bson:"residual"`
}

type Timestamps struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}

type DepositRecord struct {
	DepositId  int                 `json:"depositId" bson:"depositId"`
	CustomerId string              `json:"customerId" bson:"customerId"`
	FromDt     time.Time           `json:"fromDt" bson:"fromDt"`
	ToDt       time.Time           `json:"toDt" bson:"toDt"`
	Timestamps Timestamps          `json:"timestamps" bson:"timestamps"`
	TimeZone   string              `json:"timeZone,omitempty" bson:"timeZone,omitempty"`
	Deposit    map[int]DepositItem `json:"deposit" bson:"deposit"`
}

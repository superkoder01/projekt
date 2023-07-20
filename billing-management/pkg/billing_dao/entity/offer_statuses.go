package entity

const (
	OFFER_STATUSES = "OFFER_STATUSES"
)

type OfferStatuses struct {
	Name string `gorm:"column:NAME;primaryKey;size:45"`
}

func NewOfferStatuses() *OfferStatuses {
	return &OfferStatuses{}
}

func (tt *OfferStatuses) TableName() string {
	return OFFER_STATUSES
}

package email

import "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/offer"

type Email struct {
	Destination []string    `json:"destination"`
	Title       string      `json:"title"`
	Body        string      `json:"body"`
	Offer       offer.Offer `json:"offer"`
}

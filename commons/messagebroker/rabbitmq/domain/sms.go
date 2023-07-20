package domain

type Sms struct {
	Msisdn []string `json:"Msisdn"`
	Text   string   `json:"Text"`
}

package worker

import (
	"fmt"
	"time"
)

type Worker struct {
	ID                   int       `json:"id,omitempty"`
	FirstName            string    `json:"firstName,omitempty"`
	LastName             string    `json:"lastName,omitempty"`
	Email                string    `json:"email,omitempty"`
	Phone                string    `json:"phone,omitempty"`
	WorkStartDate        time.Time `json:"workStartDate"`
	WorkEndDate          time.Time `json:"workEndDate"`
	BlockchainAccAddress string    `json:"blockchainAccAddress,omitempty"`
	Street               string    `json:"street,omitempty"`
	City                 string    `json:"city,omitempty"`
	PostalCode           string    `json:"postalCode,omitempty"`
	Province             string    `json:"province,omitempty"`
	BuildingNumber       string    `json:"buildingNumber,omitempty"`
	ApartmentNumber      string    `json:"apartmentNumber,omitempty"`
	Country              string    `json:"country,omitempty"`
	ProviderID           int       `json:"providerId,omitempty"`
	Supervisor           int       `json:"supervisor,omitempty"`
	Role                 string    `json:"role,omitempty"`
	Status               bool      `json:"status"`
	NIP                  string    `json:"nip,omitempty"`
	REGON                string    `json:"regon,omitempty"`
	PESEL                string    `json:"pesel,omitempty"`
	KRS                  string    `json:"krs,omitempty"`
	ExtraInfo            string    `json:"extraInfo,omitempty"`
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) String() string {
	return fmt.Sprintf("%s", *w)
}

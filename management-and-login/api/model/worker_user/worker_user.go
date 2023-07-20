package worker_user

import (
	"fmt"
	"time"
)

type WorkerUser struct {
	WorkerID             int       `json:"workerId,omitempty"`
	FirstName            string    `json:"firstName,omitempty"`
	LastName             string    `json:"lastName,omitempty"`
	Email                string    `json:"email"`
	Phone                string    `json:"phone,omitempty"`
	WorkStartDate        time.Time `json:"workStartDate"`
	WorkEndDate          time.Time `json:"workEndDate"`
	BlockchainAccAddress string    `json:"blockchainAccAddress"`
	Street               string    `json:"street,omitempty"`
	City                 string    `json:"city,omitempty"`
	PostalCode           string    `json:"postalCode,omitempty"`
	Province             string    `json:"province,omitempty"`
	BuildingNumber       string    `json:"buildingNumber,omitempty"`
	ApartmentNumber      string    `json:"apartmentNumber,omitempty"`
	Country              string    `json:"country,omitempty"`
	ProviderID           int       `json:"providerId"`
	Supervisor           int       `json:"supervisor,omitempty"`
	NIP                  string    `json:"nip,omitempty"`
	REGON                string    `json:"regon,omitempty"`
	PESEL                string    `json:"pesel,omitempty"`
	KRS                  string    `json:"krs,omitempty"`
	ExtraInfo            string    `json:"extraInfo,omitempty"`
	UserID               int       `json:"userId,omitempty"`
	Login                *string   `json:"login"`
	Password             *string   `json:"password"`
	RoleID               int       `json:"roleId,omitempty"`
	Active               bool      `json:"active,omitempty"`
	MustChangePassword   bool      `json:"mustChangePassword,omitempty"`
	Status               bool      `json:"status,omitempty"`
}

func NewWorkerUser() *WorkerUser {
	return &WorkerUser{}
}

func (w *WorkerUser) String() string {
	return fmt.Sprintf("%s", *w)
}

package service_access_point

import (
	"fmt"
)

type ServiceAccessPoint struct {
	ID          int    `json:"id,omitempty"`
	SapCode     string `json:"sapCode,omitempty"`
	MeterNumber string `json:"meterNumber,omitempty"`
	City        string `json:"city,omitempty"`
	Address     string `json:"address,omitempty"`
	AccountID   int    `json:"accountId"`
	ProviderID  int    `json:"providerId"`
	Name        string `json:"name,omitempty"`
}

func NewServiceAccessPoint() *ServiceAccessPoint {
	return &ServiceAccessPoint{}
}

func (s *ServiceAccessPoint) String() string {
	return fmt.Sprintf("%s", *s)
}

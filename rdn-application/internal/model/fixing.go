package model

import "fmt"

type FixingType string

const (
	FIXING1 FixingType = "Fixing I"
	FIXING2 FixingType = "Fixing II"
)

type Fixing struct {
	Type      string  `json:"Type"`
	Date      string  `json:"Date"`
	HoursData []Model `json:"HoursData"`
}

func NewFixing(fType FixingType, date string, slice []Model) *Fixing {
	return &Fixing{
		Type:      string(fType),
		Date:      date,
		HoursData: slice,
	}
}

func (fixing *Fixing) ToString() string {
	return "" + fixing.Type + ", Date: " + fixing.Date + ", HourData: [" + fmt.Sprintf("%v", fixing.HoursData) + "]"
}

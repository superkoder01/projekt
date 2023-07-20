package entity

import (
	"fmt"
)

type HourData struct {
	Hour  string
	Value float64
}

func (data HourData) GetValue() float64 {
	return data.Value
}

func (data HourData) GetHour() string {
	return data.Hour
}

func (data HourData) ToString() string {
	return "hour: " + fmt.Sprintf("%v", data.Hour) + ", value: " + fmt.Sprintf("%v", data.Value)
}

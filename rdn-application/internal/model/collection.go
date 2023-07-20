package model

import "time"

type Collection interface {
	AppendData(key string, value float64)
	AppendDate(date time.Time)
	GetCollectionData() (map[string]float64, time.Time)
}

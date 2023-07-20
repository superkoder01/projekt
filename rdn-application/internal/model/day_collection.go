package model

import (
	"fmt"
	"time"
)

type DayCollection struct {
	Date     time.Time
	HourData map[string]float64
}

func (collection *DayCollection) GetCollectionData() (map[string]float64, time.Time) {
	return collection.HourData, collection.Date
}

func (collection *DayCollection) AppendDate(date time.Time) {
	collection.Date = date
}

func (collection *DayCollection) AppendData(key string, value float64) {
	collection.HourData[key] = value
}

func (collection *DayCollection) ToString() string {
	return fmt.Sprintf("%s", *collection)
}

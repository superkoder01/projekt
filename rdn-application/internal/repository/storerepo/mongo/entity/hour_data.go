package entity

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HourData struct {
	Date  primitive.DateTime `bson:"timestamp"`
	Value float32            `bson:"value"`
	Time  int                `bson:"time"`
}

func (data *HourData) GetData() time.Time {
	return data.Date.Time()
}

func (data *HourData) GetValue() float32 {
	return data.Value
}

func (data *HourData) GetHour() int {
	return data.Time
}

func (data *HourData) ToString() string {
	return fmt.Sprintf("HourData: {%v %v %v}", data.Date, data.Time, data.Value)
}

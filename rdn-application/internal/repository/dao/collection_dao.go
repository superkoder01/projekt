package dao

type ModelDao interface {
	ToString() string
}

type HourDataDao interface {
	GetValue() float64
	GetHour() string
	ModelDao
}

type DailyDataDao interface {
	GetHourData() []HourDataDao
	GetDate() string
	AppendHourData(dao interface{})
	ModelDao
}

package entity

import (
	"RDN-application/internal/repository/dao"
	"fmt"
)

type DailyData struct {
	Hours []HourData
	Date  string
}

func (daily *DailyData) GetHourData() []dao.HourDataDao {
	var daos []dao.HourDataDao
	for _, hour := range daily.Hours {
		daos = append(daos, hour)
	}
	return daos
}

func (daily *DailyData) AppendHourData(dao interface{}) {
	daily.Hours = append(daily.Hours, dao.(HourData))
}

func (daily *DailyData) GetDate() string {
	return daily.Date
}

func (daily *DailyData) ToString() string {
	return "day: " + fmt.Sprintf("%v", daily.Hours) + ", date: " + fmt.Sprintf("%v", daily.Date)
}

package entity

import (
	"RDN-application/internal/model"
	"RDN-application/internal/repository/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
	"time"
)

type entityMapper struct{}

func (entityMapper) MapDayCollectionDtoToDao(slice model.Collection, date time.Time) []dao.HourDataDao {
	var daos []dao.HourDataDao

	for key, value := range slice.GetCollectionData() {
		daos = append(daos, &HourData{
			Date:  primitive.NewDateTimeFromTime(date),
			Value: value,
			Time:  key,
		})
	}
	sort.Slice(daos, func(i, j int) bool {
		return daos[i].GetHour() < daos[j].GetHour()
	})
	return daos
}

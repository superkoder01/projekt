package entity

import (
	"RDN-application/internal/model"
	"RDN-application/internal/repository/dao"
)

func MapDayCollectionDtoToDao(slice model.Collection) []dao.HourDataDao {
	var data []dao.HourDataDao

	collectionData, _ := slice.GetCollectionData()
	for key, value := range collectionData {
		data = append(data, &HourData{
			Hour:  key,
			Value: value,
		})
	}
	return data
}

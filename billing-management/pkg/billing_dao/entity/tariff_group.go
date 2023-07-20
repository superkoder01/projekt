package entity

import "time"

const (
	TARIFF_GROUP = "TARIFF_GROUP"
)

type TariffGroupEntity interface {
	SetDistributionNetworkOperatorID(int)
	SetTariffGroupLabelName(int)
	SetName(string)
	SetStartDate(time.Time)
	SetEndDate(time.Time)
}

type TariffGroup struct {
	ID                            int       `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	DistributionNetworkOperatorID int       `gorm:"column:DISTRIBUTION_NETWORK_OPERATOR_ID;size:11;not null"`
	TariffGroupLabelName          string    `gorm:"column:TARIFF_GROUP_LABEL_NAME;size:11;not null"`
	Name                          string    `gorm:"column:NAME;size:45;default:null"`
	StartDate                     time.Time `gorm:"column:START_DATE"`
	EndDate                       time.Time `gorm:"column:END_DATE"`
}

func NewTariffGroup() *TariffGroup {
	return &TariffGroup{}
}

func (tg *TariffGroup) TableName() string {
	return TARIFF_GROUP
}

func (tg *TariffGroup) SetDistributionNetworkOperatorID(i int) {
	tg.DistributionNetworkOperatorID = i
}

func (tg *TariffGroup) SetTariffGroupLabelName(i string) {
	tg.TariffGroupLabelName = i
}

func (tg *TariffGroup) SetName(s string) {
	tg.Name = s
}

func (tg *TariffGroup) SetStartDate(t time.Time) {
	tg.StartDate = t
}

func (tg *TariffGroup) SetEndDate(t time.Time) {
	tg.EndDate = t
}

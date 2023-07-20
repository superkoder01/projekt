package entity

const (
	TARIFF_GROUP_PARAMETER = "TARIFF_GROUP_PARAMETER"
)

type TariffGroupParameterEntity interface {
	TariffGroupID(int)
	SetParameterNameID(int)
	SetPrice(float64)
}

type TariffGroupParameter struct {
	ID              int            `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	TariffGroupID   int            `gorm:"column:TARIFF_GROUP_ID;size:11;not null"`
	ParameterNameID int            `gorm:"column:PARAMETER_NAME_ID;size:11;not null"`
	ParameterName   *ParameterName `gorm:"foreignKey:PARAMETER_NAME_ID;references:ID"`
	Price           float64        `gorm:"column:PRICE;size:45;default:null"`
}

func NewTariffGroupParameter() *TariffGroupParameter {
	return &TariffGroupParameter{}
}

func (tgp *TariffGroupParameter) TableName() string {
	return TARIFF_GROUP_PARAMETER
}

func (tgp *TariffGroupParameter) SetTariffGroupID(i int) {
	tgp.TariffGroupID = i
}

func (tgp *TariffGroupParameter) SetParameterNameID(i int) {
	tgp.ParameterNameID = i
}

func (tgp *TariffGroupParameter) SetPrice(s float64) {
	tgp.Price = s
}

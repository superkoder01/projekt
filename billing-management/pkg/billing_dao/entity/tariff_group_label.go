package entity

const (
	TARIFF_GROUP_LABEL = "TARIFF_GROUP_LABEL"
)

type TariffGroupLabelEntity interface {
	SetName(string)
}
type TariffGroupLabel struct {
	Name string `gorm:"column:NAME;size:45;primaryKey"`
}

func NewTariffGroupLabel() *TariffGroupLabel {
	return &TariffGroupLabel{}
}

func (pn *TariffGroupLabel) TableName() string {
	return TARIFF_GROUP_LABEL
}

func (pn *TariffGroupLabel) SetName(s string) {
	pn.Name = s
}

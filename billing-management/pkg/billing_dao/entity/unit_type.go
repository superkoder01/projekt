package entity

type UnitTypeName string

const (
	UNIT_TYPE = "UNIT_TYPE"

	SECONDS                UnitTypeName = "SECONDS"
	TOTAL_OCTETS           UnitTypeName = "TOTAL_OCTETS"
	INPUT_OCTETS           UnitTypeName = "INPUT_OCTETS"
	OUTPUT_OCTETS          UnitTypeName = "OUTPUT_OCTETS"
	SERVICE_SPECIFIC_UNITS UnitTypeName = "SERVICE_SPECIFIC_UNITS"
)

type UnitType struct {
	ID   int          `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name UnitTypeName `gorm:"column:NAME" sql:"type:ENUM('SECONDS','TOTAL_OCTETS','INPUT_OCTETS','OUTPUT_OCTETS','SERVICE_SPECIFIC_UNITS')"`
}

func NewUnitType() *UnitType {
	return &UnitType{}
}

func (tt *UnitType) TableName() string {
	return UNIT_TYPE
}

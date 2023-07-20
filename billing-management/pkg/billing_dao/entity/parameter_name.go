package entity

const (
	PARAMETER_NAME = "PARAMETER_NAME"
)

type ParameterNameEntity interface {
	SetName(string)
	SetCode(string)
}

type ParameterName struct {
	ID   int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name string `gorm:"column:NAME;size:45;not null"`
	Code string `gorm:"column:CODE;size:45;default:null"`
}

func NewParameterName() *ParameterName {
	return &ParameterName{}
}

func (pn *ParameterName) TableName() string {
	return PARAMETER_NAME
}

func (pn *ParameterName) SetName(s string) {
	pn.Name = s
}

func (pn *ParameterName) SetCode(s string) {
	pn.Code = s
}

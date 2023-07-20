package entity

type SettingType string

const (
	SETTINGS = "SETTINGS"

	INTEGER SettingType = "INTEGER"
	BOOLEAN SettingType = "BOOLEAN"
	FLOAT   SettingType = "FLOAT"
)

// TODO: error handling
func (c SettingType) Value() string {
	return string(c)
}

type SettingsEntity interface {
	SetName(string)
	SetValue(string)
	SetType(SettingType)
}

type Settings struct {
	ID    int         `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name  string      `gorm:"column:NAME;size:45;default:null"`
	Value string      `gorm:"column:VALUE;size:45;default:null"`
	Type  SettingType `gorm:"column:TYPE" sql:"type:ENUM('INTEGER','BOOLEAN','FLOAT')"`
}

func NewSettings() *Settings {
	return &Settings{}
}

func (z *Settings) TableName() string {
	return ZONE
}

func (z *Settings) SetName(s string) {
	z.Name = s
}

func (z *Settings) SetValue(s string) {
	z.Value = s
}

func (z *Settings) SetType(s SettingType) {
	z.Type = s
}

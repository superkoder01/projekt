package enum

type SettingType string

const (
	INTEGER SettingType = "INTEGER"
	BOOLEAN SettingType = "BOOLEAN"
	FLOAT   SettingType = "FLOAT"
)

func (s SettingType) Name() string {
	switch s {
	case INTEGER:
		return "INTEGER"
	case BOOLEAN:
		return "BOOLEAN"
	case FLOAT:
		return "FLOAT"
	default:
		return ""
	}
}

package entity

import "database/sql/driver"

type BusinessTypeName string

const (
	BT_B2B BusinessTypeName = "B2B"
	BT_B2C BusinessTypeName = "B2C"
)

func (b *BusinessTypeName) Scan(value interface{}) error {
	*b = BusinessTypeName(value.([]byte))
	return nil
}

func (b BusinessTypeName) Value() (driver.Value, error) {
	return string(b), nil
}

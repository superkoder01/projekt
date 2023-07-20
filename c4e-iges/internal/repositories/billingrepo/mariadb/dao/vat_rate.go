package mariadb_dao

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type vatGroup string

const (
	Gxx vatGroup = "Gxx"
)

func (vg *vatGroup) Scan(value interface{}) error {
	*vg = vatGroup(value.([]byte))
	return nil
}

func (vg vatGroup) Value() (driver.Value, error) {
	return string(vg), nil
}

type VatRate struct {
	Value int `gorm:"column:RATE"`
}

func (VatRate) TableName() string {
	return "VAT_RATES"
}

func (v VatRate) String() string {
	//data, _ := json.Marshal(p)
	data, _ := json.MarshalIndent(v, "", "\t")
	return fmt.Sprintf("%s", data)
}

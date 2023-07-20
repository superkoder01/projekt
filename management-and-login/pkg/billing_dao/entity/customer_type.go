package entity

import "database/sql/driver"

type CustomerTypeName string

const (
	// table name
	CUSTOMER_TYPE = "CUSTOMER_TYPE"

	CT_CONSUMER CustomerTypeName = "CONSUMER"
	CT_PROSUMER CustomerTypeName = "PROSUMER"
	CT_PRODUCER CustomerTypeName = "PRODUCER"
)

func (c *CustomerTypeName) Scan(value interface{}) error {
	*c = CustomerTypeName(value.([]byte))
	return nil
}

func (c CustomerTypeName) Value() (driver.Value, error) {
	return string(c), nil
}

type CustomerType struct {
	ID   int              `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name CustomerTypeName `gorm:"column:NAME" sql:"type:ENUM('CONSUMER', 'PROSUMER','PRODUCER')"`
}

func (c *CustomerType) TableName() string {
	return CUSTOMER_TYPE
}

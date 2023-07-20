package entity

import "database/sql/driver"

type RoleName string

const (
	// table name
	ROLE = "ROLE"

	SUPER_ADMIN         RoleName = "SUPER_ADMIN"
	ADMINISTRATOR_FULL  RoleName = "ADMINISTRATOR_FULL"
	ADMINISTRATOR_BASIC RoleName = "ADMINISTRATOR_BASIC"
	TRADER              RoleName = "TRADER"
	SUPER_AGENT         RoleName = "SUPER_AGENT"
	AGENT               RoleName = "AGENT"
	PROSUMER            RoleName = "PROSUMER"
	ACCOUNTER           RoleName = "ACCOUNTER"
)

func (r *RoleName) Scan(value interface{}) error {
	*r = RoleName(value.([]byte))
	return nil
}

func (r RoleName) Value() (driver.Value, error) {
	return string(r), nil
}

type Role struct {
	ID   int      `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name RoleName `gorm:"column:NAME" sql:"type:ENUM('SUPER_ADMIN', 'ADMINISTRATOR_FULL','ADMINISTRATOR_BASIC','TRADER','SUPER_AGENT','AGENT','PROSUMER', 'ACCOUNTER')"`
}

func (r *Role) TableName() string {
	return ROLE
}

package mariadb_dao

import "database/sql/driver"

type invoiceStatus string

const (
	ISSUED invoiceStatus = "ISSUED"
	PAID   invoiceStatus = "PAID"
)

func (ct *invoiceStatus) Scan(value interface{}) error {
	*ct = invoiceStatus(value.([]byte))
	return nil
}

func (ct invoiceStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Invoice struct {
	InvoiceNumber string        `gorm:"column:INVOICE_NUMBER"`
	ContractId    uint          `gorm:"column:CONTRACT_ID"`
	InvoiceStatus invoiceStatus `gorm:"column:STATUS"`
}

func (Invoice) TableName() string {
	return "INVOICE"
}

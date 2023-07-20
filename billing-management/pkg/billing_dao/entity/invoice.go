package entity

type InvoiceStatus string

const (
	INVOICE = "INVOICE"

	ISSUED InvoiceStatus = "ISSUED"
	PAID   InvoiceStatus = "PAID"
)

type InvoiceEntity interface {
	SetContractID(int)
	SetInvoiceNumber(string)
	SetStatus(InvoiceStatus)
}

// TODO: error handling
func (c InvoiceStatus) Value() string {
	return string(c)
}

type Invoice struct {
	ID            int           `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ContractID    int           `gorm:"column:CONTRACT_ID;size:11;not null"`
	InvoiceNumber string        `gorm:"column:INVOICE_NUMER;size:55;default:null"`
	Status        InvoiceStatus `gorm:"column:STATUS" sql:"type:ENUM('ISSUED','PAID')"`
}

func NewInvoice() *Invoice {
	return &Invoice{}
}

func (inv *Invoice) TableName() string {
	return INVOICE
}

func (inv *Invoice) SetContractID(i int) {
	inv.ContractID = i
}

func (inv *Invoice) SetInvoiceNumber(s string) {
	inv.InvoiceNumber = s
}

func (inv *Invoice) SetStatus(invs InvoiceStatus) {
	inv.Status = invs
}

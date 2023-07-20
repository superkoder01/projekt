package enum

type InvoiceStatus string

const (
	ISSUED InvoiceStatus = "ISSUED"
	PAID   InvoiceStatus = "PAID"
)

func (s InvoiceStatus) Name() string {
	switch s {
	case ISSUED:
		return "ISSUED"
	case PAID:
		return "PAID"
	default:
		return ""
	}
}

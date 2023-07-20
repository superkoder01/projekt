package enum

type CustomerType int

const (
	CT_CONSUMER CustomerType = iota + 1
	CT_PROSUMER
	CT_PRODUCER
)

func (c CustomerType) Name() string {
	switch c {
	case CT_CONSUMER:
		return "CONSUMER"
	case CT_PROSUMER:
		return "PROSUMER"
	case CT_PRODUCER:
		return "PRODUCER"
	default:
		return ""
	}
}

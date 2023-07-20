package enum

type ContractStatus string

const (
	CS_DRAFT     ContractStatus = "DRAFT"
	CS_FINAL     ContractStatus = "FINAL"
	CS_SENT      ContractStatus = "SENT"
	CS_DELIVERED ContractStatus = "DELIVERED"
	CS_ACCEPTED  ContractStatus = "ACCEPTED"
)

func (s ContractStatus) Name() string {
	switch s {
	case CS_DRAFT:
		return "DRAFT"
	case CS_FINAL:
		return "FINAL"
	case CS_SENT:
		return "SENT"
	case CS_DELIVERED:
		return "DELIVERED"
	case CS_ACCEPTED:
		return "ACCEPTED"
	default:
		return ""
	}
}

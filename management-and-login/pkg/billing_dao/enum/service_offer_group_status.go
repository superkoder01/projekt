package enum

type ServiceOfferGroupStatus string

const (
	SOGS_DRAFT       ServiceOfferGroupStatus = "DRAFT"
	SOGS_FINAL       ServiceOfferGroupStatus = "FINAL"
	SOGS_SENT        ServiceOfferGroupStatus = "SENT"
	SOGS_DELIVERED   ServiceOfferGroupStatus = "DELIVERED"
	SOGS_NEGOTIATION ServiceOfferGroupStatus = "NEGOTIATION"
)

func (s ServiceOfferGroupStatus) Name() string {
	switch s {
	case SOGS_DRAFT:
		return "DRAFT"
	case SOGS_FINAL:
		return "FINAL"
	case SOGS_SENT:
		return "SENT"
	case SOGS_DELIVERED:
		return "DELIVERED"
	case SOGS_NEGOTIATION:
		return "NEGOTIATION"
	default:
		return ""
	}
}

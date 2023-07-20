package enum

type ServiceType string

const (
	SALE       ServiceType = "SALE"
	REPURCHASE ServiceType = "REPURCHASE"
)

func (c ServiceType) Name() string {
	switch c {
	case SALE:
		return "SALE"
	case REPURCHASE:
		return "REPURCHASE"
	default:
		return ""
	}
}

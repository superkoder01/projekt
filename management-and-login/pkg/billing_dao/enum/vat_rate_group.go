package enum

type VatRateGroup string

const (
	Gxx       VatRateGroup = "Gxx"
	OWN_USAGE VatRateGroup = "OWN_USAGE"
	RESALE    VatRateGroup = "RESALE"
)

func (s VatRateGroup) Name() string {
	switch s {
	case Gxx:
		return "Gxx"
	case OWN_USAGE:
		return "OWN_USAGE"
	case RESALE:
		return "RESALE"
	default:
		return ""
	}
}

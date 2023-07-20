package enum

type RatingPlanTypeName string

const (
	RPT_SALE       RatingPlanTypeName = "SALE"
	RPT_REPURCHASE RatingPlanTypeName = "REPURCHASE"
	RPT_RDN        RatingPlanTypeName = "RDN"
)

func (s RatingPlanTypeName) Name() string {
	switch s {
	case RPT_SALE:
		return "SALE"
	case RPT_REPURCHASE:
		return "REPURCHASE"
	case RPT_RDN:
		return "RDN"
	default:
		return ""
	}
}

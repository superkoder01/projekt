package entity

type RatingPlanTypeName string

const (
	RATING_PLAN_TYPE = "RATING_PLAN_TYPE"

	RPT_SALE       RatingPlanTypeName = "SALE"
	RPT_REPURCHASE RatingPlanTypeName = "REPURCHASE"
	RPT_RDN        RatingPlanTypeName = "RDN"
)

// TODO: error handling
func (c RatingPlanTypeName) Value() string {
	return string(c)
}

type RatingPlanType struct {
	ID   int                `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	Name RatingPlanTypeName `gorm:"column:NAME" sql:"type:ENUM('SALE', 'REPURCHASE', 'RDN')"`
}

func NewRatingPlanType() *RatingPlanType {
	return &RatingPlanType{}
}

func (rpt *RatingPlanType) TableName() string {
	return RATING_PLAN_TYPE
}

package entity

const (
	REGION = "REGION"
)

type RegionEntity interface {
	SetProviderID(int)
	SetName(string)
}

type Region struct {
	ID                            int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	DistributionNetworkOperatorID int    `gorm:"column:DISTRIBUTION_NETWORK_OPERATOR_ID;size:11;not null"`
	Name                          string `gorm:"column:NAME;size:45;default:null"`
}

func NewRegion() *Region {
	return &Region{}
}

func (r *Region) TableName() string {
	return REGION
}

func (r *Region) SetDistributionNetworkOperatorID(i int) {
	r.DistributionNetworkOperatorID = i
}

func (r *Region) SetName(s string) {
	r.Name = s
}

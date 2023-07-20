package entity

type Units string

const (
	ZONE = "ZONE"

	kWh Units = "kWh"
	mWh Units = "mWh"
)

// TODO: error handling
func (c Units) Value() string {
	return string(c)
}

type ZoneEntity interface {
	SetProviderID(int)
	SetRatingPlanID(int)
	SetDescription(string)
	SetHours(string)
	SetNetPrice(float64)
	SetUnits(Units)
}

type Zone struct {
	ID           int     `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	RatingPlanID int     `gorm:"column:RATING_PLAN_ID;size:11;not null"`
	Description  string  `gorm:"column:DESCRIPTION;size:100;default:null"`
	Hours        string  `gorm:"column:HOURS;size:45;default:null"`
	NetPrice     float64 `gorm:"column:NET_PRICE;"`
	Units        Units   `gorm:"column:UNITS" sql:"type:ENUM('kWh','mWh')"`
}

func NewZone() *Zone {
	return &Zone{}
}

func (z *Zone) TableName() string {
	return ZONE
}

func (z *Zone) SetRatingPlanID(i int) {
	z.RatingPlanID = i
}

func (z *Zone) SetDescription(s string) {
	z.Description = s
}

func (z *Zone) SetHours(s string) {
	z.Hours = s
}

func (z *Zone) SetNetPrice(s float64) {
	z.NetPrice = s
}

func (z *Zone) SetUnits(s Units) {
	z.Units = s
}

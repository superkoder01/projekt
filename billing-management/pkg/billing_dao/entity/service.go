package entity

type ServiceType string

const (
	SERVICE = "SERVICE"

	SALE       ServiceType = "SALE"
	REPURCHASE ServiceType = "REPURCHASE"
)

// TODO: error handling
func (s ServiceType) Value() string {
	return string(s)
}

type ServiceEntity interface {
	SetProviderID(int)
	SetName(string)
	SetType(ServiceType)
}

type Service struct {
	ID         int         `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	ProviderID int         `gorm:"column:PROVIDER_ID;size:11;not null"`
	Name       string      `gorm:"column:NAME;size:45;default:null"`
	Type       ServiceType `gorm:"column:TYPE" sql:"type:ENUM('SALE', 'REPURCHASE')"`
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) TableName() string {
	return SERVICE
}

func (s *Service) SetProviderID(i int) {
	s.ProviderID = i
}

func (s *Service) SetName(name string) {
	s.Name = name
}

func (s *Service) SetType(t ServiceType) {
	s.Type = t
}

package entity

const (
	SERVICE_ACCESS_POINT = "SERVICE_ACCESS_POINT"
)

type ServiceAccessPointEntity interface {
	SetAccountID(int)
	SetCity(string)
	SetMeterNumber(string)
	SetAddress(string)
	SetSapCode(string)
	SetProviderID(int)
	SetName(string)
}

type ServiceAccessPoint struct {
	ID          int    `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	AccountID   int    `gorm:"column:ACCOUNT_ID;size:11;not null"`
	City        string `gorm:"column:CITY;size:45;default:null"`
	Address     string `gorm:"column:ADDRESS;size:45;default:null"`
	MeterNumber string `gorm:"column:METER_NUMBER;size:45;default:null"`
	SapCode     string `gorm:"column:SAP_CODE;size:45;default:null"`
	ProviderID  int    `gorm:"column:PROVIDER_ID;size:11;not null"`
	Name        string `gorm:"column:NAME;size:45;default:null"`
}

func NewServiceAccessPoint() *ServiceAccessPoint {
	return &ServiceAccessPoint{}
}

func (sap *ServiceAccessPoint) TableName() string {
	return SERVICE_ACCESS_POINT
}

func (sap *ServiceAccessPoint) SetAccountID(i int) {
	sap.AccountID = i
}

func (sap *ServiceAccessPoint) SetCity(s string) {
	sap.City = s
}

func (sap *ServiceAccessPoint) SetMeterNumber(s string) {
	sap.MeterNumber = s
}

func (sap *ServiceAccessPoint) SetAddress(s string) {
	sap.Address = s
}

func (sap *ServiceAccessPoint) SetSapCode(s string) {
	sap.SapCode = s
}

func (sap *ServiceAccessPoint) SetProviderID(i int) {
	sap.ProviderID = i
}

func (sap *ServiceAccessPoint) SetName(s string) {
	sap.Name = s
}

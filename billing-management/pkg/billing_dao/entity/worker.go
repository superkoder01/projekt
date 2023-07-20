package entity

import "time"

const (
	WORKER = "WORKER"
)

type WorkerEntity interface {
	SetFirstName(string)
	SetLastName(string)
	SetEmail(string)
	SetPhone(string)
	SetWorkStartDate(time.Time)
	SetWorkEndDate(time.Time)
	SetBlockchainAccAddress(string)
	SetCity(string)
	SetPostalCode(string)
	SetBuildingNumber(string)
	SetApartmentNumber(string)
	SetCountry(string)
	SetProviderID(int)
	SetSupervisor(int)
	SetStatus(bool)
	SetNIP(string)
	SetREGON(string)
	SetKRS(string)
	SetPESEL(string)
	SetExtraInfo(string)
}

type Worker struct {
	ID                   int       `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	FirstName            string    `gorm:"column:FIRST_NAME;size:45;default:null"`
	LastName             string    `gorm:"column:LAST_NAME;size:45;default:null"`
	Email                string    `gorm:"column:EMAIL;size:45;not null;unique"`
	Phone                string    `gorm:"column:PHONE;size:45;default:null"`
	WorkStartDate        time.Time `gorm:"column:WORK_START_DATE;default:null"`
	WorkEndDate          time.Time `gorm:"column:WORK_END_DATE;default:null"`
	BlockchainAccAddress string    `gorm:"column:BLOCKCHAIN_ACC_ADDRESS;size:75;default:null;unique"`
	Street               string    `gorm:"column:STREET;size:45;default:null"`
	City                 string    `gorm:"column:CITY;size:45;default:null"`
	PostalCode           string    `gorm:"column:POSTAL_CODE;size:45;default:null"`
	Province             string    `gorm:"column:PROVINCE;size:45;default:null"`
	BuildingNumber       string    `gorm:"column:BUILDING_NUMBER;size:45;default:null"`
	ApartmentNumber      string    `gorm:"column:APARTMENT_NUMBER;size:45;default:null"`
	Country              string    `gorm:"column:COUNTRY;size:45;default:null"`
	ProviderID           int       `gorm:"column:PROVIDER_ID;size:11;not null"`
	Supervisor           int       `gorm:"column:SUPERVISOR;size:11;default:null"`
	Worker               *Worker   `gorm:"foreignKey:SUPERVISOR;references:ID"`
	Status               bool      `gorm:"column:STATUS;default:0"`
	NIP                  string    `gorm:"column:NIP;size:45;default:null"`
	REGON                string    `gorm:"column:REGON;size:45;default:null"`
	PESEL                string    `gorm:"column:PESEL;size:45;default:null"`
	KRS                  string    `gorm:"column:KRS;size:45;default:null"`
	ExtraInfo            string    `gorm:"column:EXTRA_INFO;size:45;default:null"`
}

type WorkerJoinUserRole struct {
	ID                   int       `gorm:"column:ID;size:11;primaryKey;autoIncrement"`
	FirstName            string    `gorm:"column:FIRST_NAME;size:45;default:null"`
	LastName             string    `gorm:"column:LAST_NAME;size:45;default:null"`
	Email                string    `gorm:"column:EMAIL;size:45;not null;unique"`
	Phone                string    `gorm:"column:PHONE;size:45;default:null"`
	WorkStartDate        time.Time `gorm:"column:WORK_START_DATE;default:null"`
	WorkEndDate          time.Time `gorm:"column:WORK_END_DATE;default:null"`
	BlockchainAccAddress string    `gorm:"column:BLOCKCHAIN_ACC_ADDRESS;size:75;default:null;unique"`
	Street               string    `gorm:"column:STREET;size:45;default:null"`
	City                 string    `gorm:"column:CITY;size:45;default:null"`
	PostalCode           string    `gorm:"column:POSTAL_CODE;size:45;default:null"`
	Province             string    `gorm:"column:PROVINCE;size:45;default:null"`
	BuildingNumber       string    `gorm:"column:BUILDING_NUMBER;size:45;default:null"`
	ApartmentNumber      string    `gorm:"column:APARTMENT_NUMBER;size:45;default:null"`
	Country              string    `gorm:"column:COUNTRY;size:45;default:null"`
	ProviderID           int       `gorm:"column:PROVIDER_ID;size:11;not null"`
	Supervisor           int       `gorm:"column:SUPERVISOR;size:11;default:null"`
	Worker               *Worker   `gorm:"foreignKey:SUPERVISOR;references:ID"`
	Status               bool      `gorm:"column:STATUS;default:0"`
	RoleID               int       `gorm:"column:ROLE_ID;size:11;default:null"`
	NIP                  string    `gorm:"column:NIP;size:45;default:null"`
	REGON                string    `gorm:"column:REGON;size:45;default:null"`
	PESEL                string    `gorm:"column:PESEL;size:45;default:null"`
	KRS                  string    `gorm:"column:KRS;size:45;default:null"`
	ExtraInfo            string    `gorm:"column:EXTRA_INFO;size:45;default:null"`
}

func NewWorkerJoinUserRole() *WorkerJoinUserRole {
	return &WorkerJoinUserRole{}
}

func (w *WorkerJoinUserRole) TableName() string {
	return WORKER
}

func NewWorker() *Worker {
	return &Worker{}
}

func (w *Worker) TableName() string {
	return WORKER
}

func (w *Worker) SetFirstName(s string) {
	w.FirstName = s
}

func (w *Worker) SetLastName(s string) {
	w.LastName = s
}

func (w *Worker) SetEmail(s string) {
	w.Email = s
}

func (w *Worker) SetPhone(s string) {
	w.Phone = s
}

func (w *Worker) SetWorkStartDate(t time.Time) {
	w.WorkStartDate = t
}

func (w *Worker) SetWorkEndDate(t time.Time) {
	w.WorkEndDate = t
}

func (w *Worker) SetBlockchainAccAddress(s string) {
	w.BlockchainAccAddress = s
}

func (w *Worker) SetCity(s string) {
	w.City = s
}

func (w *Worker) SetPostalCode(s string) {
	w.PostalCode = s
}

func (w *Worker) SetBuildingNumber(s string) {
	w.BuildingNumber = s
}

func (w *Worker) SetApartmentNumber(s string) {
	w.ApartmentNumber = s
}

func (w *Worker) SetCountry(s string) {
	w.Country = s
}

func (w *Worker) SetStatus(s bool) {
	w.Status = s
}

func (w *Worker) SetProviderID(i int) {
	w.ProviderID = i
}

func (w *Worker) SetSupervisor(i int) {
	w.Supervisor = i
}

func (w *Worker) SetNIP(s string) {
	w.NIP = s
}

func (w *Worker) SetKRS(s string) {
	w.KRS = s
}

func (w *Worker) SetREGON(s string) {
	w.REGON = s
}

func (w *Worker) SetPESEL(s string) {
	w.PESEL = s
}

func (w *Worker) SetExtraInfo(s string) {
	w.ExtraInfo = s
}

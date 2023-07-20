package model

import (
	"time"
)

type Transaction struct {
	ID            uint `gorm:"primarykey"`
	OperationType string
	ObjectName    string
	Author        string
	CreatedAt     time.Time
	Transactions  []TransactionDetails `gorm:"many2many:transaction_transaction_details;"`
}

type TransactionDetails struct {
	ID         uint `gorm:"primarykey"`
	Info       string
	AuthorRole string
}

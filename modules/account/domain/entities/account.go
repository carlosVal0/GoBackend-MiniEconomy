package entities

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	ID             int
	AccountNumber  string
	CurrentBalance float64
	Status         bool `gorm:"default:true"`
	UserID         int
	Transactions   []Transaction
}

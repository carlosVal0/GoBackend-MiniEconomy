package entities

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID                   int
	OriginAccountNumber  string
	DestinyAccountNumber string
	Amount               float64
	AccountID            int
}

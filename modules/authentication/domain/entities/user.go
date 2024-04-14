package entities

import (
	"time"

	accountDomain "github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        int `gorm:"primaryKey"`
	Name      string
	Password  string
	Email     string
	OrgId     int
	Account   []accountDomain.Account `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

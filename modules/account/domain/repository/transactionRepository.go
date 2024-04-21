package repository

import (
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/core/database"
)

func CreateTransaction(tx *entities.Transaction) error {

	db := database.GetDbBean()
	result := db.Select("origin_account_number", "destiny_account_number", "amount", "account_id").Create(&tx)
	return result.Error

}

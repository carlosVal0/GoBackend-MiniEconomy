package repository

import (
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	authRepository "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/repository"
)

func CreateTransaction(tx *entities.Transaction) error {

	db := authRepository.GetDbBean()
	result := db.Select("origin_account_number", "destiny_account_number", "amount", "account_id").Create(&tx)
	return result.Error

}

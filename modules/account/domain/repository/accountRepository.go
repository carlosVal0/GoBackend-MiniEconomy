package repository

import (
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	authRepository "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/repository"
)

func CreateAccount(acc *entities.Account) error {

	db := authRepository.GetDbBean()
	result := db.Select("account_number", "current_balance", "user_id").Create(acc)
	return result.Error
}

func GetAccounts(userId int) ([]entities.Account, error) {
	var accounts []entities.Account
	db := authRepository.GetDbBean()
	result := db.Model(&entities.Account{}).Where("user_id = ?", userId).Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}
	return accounts, nil
}

func UpdateAccount(acc *entities.Account) error {
	db := authRepository.GetDbBean()
	result := db.Model(&acc).Select("current_balance", "status").Updates(&acc)
	return result.Error

}

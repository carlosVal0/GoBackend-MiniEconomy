package repository

import (
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/core/database"
)

func CreateAccount(acc *entities.Account) error {

	db := database.GetDbBean()
	result := db.Select("account_number", "current_balance", "user_id").Create(acc)
	return result.Error
}

func GetAccounts(userId int) ([]entities.Account, error) {
	var accounts []entities.Account
	db := database.GetDbBean()
	result := db.Model(&entities.Account{}).Where("user_id = ?", userId).Find(&accounts)
	if result.Error != nil {
		return nil, result.Error
	}
	return accounts, nil
}

func GetAccountByNumber(accNumber string) (*entities.Account, error) {

	var account entities.Account
	db := database.GetDbBean()
	result := db.Model(&entities.Account{}).Where("account_number = ?", accNumber).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil

}

func UpdateAccount(acc *entities.Account) error {
	db := database.GetDbBean()
	result := db.Model(&acc).Select("current_balance", "status").Updates(&acc)
	return result.Error

}

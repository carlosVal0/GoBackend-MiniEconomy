package application

import (
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	transactionRepository "github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/repository"
)

func ExecuteTransaction(tx *entities.Transaction, userId int) error {
	result := transactionRepository.CreateTransaction(tx)
	if result != nil {
		return result
	}

	return result
}

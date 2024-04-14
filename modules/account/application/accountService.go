package application

import (
	"math/rand"

	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/repository"
)

const (
	seqSize = 10
)

func CreateAccountService(userId int) (*entities.Account, error) {
	acc := &entities.Account{
		AccountNumber:  GenerateAccountSeq(),
		CurrentBalance: 0.0,
		UserID:         userId,
	}
	err := repository.CreateAccount(acc)
	if err != nil {
		return nil, err
	}

	return acc, nil

}

func GetAccountsService(userId int) ([]entities.Account, error) {
	accs, err := repository.GetAccounts(userId)
	if err != nil {
		return nil, err
	}

	return accs, nil
}

func GenerateAccountSeq() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, seqSize)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

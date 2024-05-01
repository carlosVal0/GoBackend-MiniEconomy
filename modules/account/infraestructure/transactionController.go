package infraestructure

import (
	"fmt"
	"net/http"

	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/application"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/repository"
	authApplication "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/application"
	"github.com/gin-gonic/gin"
)

type TransactionDTO struct {
	OriginAccountNumber  string  `json:"origin_acount_number"`
	DestinyAccountNumber string  `json:"destiny_acount_number"`
	Amount               float64 `json:"amount"`
}

func TransferFunds(ctx *gin.Context) {

	claims := TokenHandleTransaction(ctx)
	if claims == nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "empty auth header",
		})
		return
	}

	var transactionRequest TransactionDTO
	errBind := ctx.BindJSON(&transactionRequest)
	if errBind != nil {
		fmt.Println(errBind)
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Bad request",
		})
		return
	}

	accounts, err := application.GetAccountsService(claims.Id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error consulting accounts",
		})
		return
	}

	var originAccount *entities.Account
	for _, acc := range accounts {
		if acc.AccountNumber == transactionRequest.OriginAccountNumber {
			originAccount = &acc
		}
	}

	if originAccount.CurrentBalance-transactionRequest.Amount < 0 {
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"message": "Error consulting accounts",
		})
		return
	}

	destinyAccount, errDestAcc := repository.GetAccountByNumber(transactionRequest.DestinyAccountNumber)
	if errDestAcc != nil {
		ctx.IndentedJSON(http.StatusNotAcceptable, gin.H{
			"message": "Error consulting destiny accounts",
		})
		return
	}

	originAccount.CurrentBalance -= transactionRequest.Amount
	destinyAccount.CurrentBalance += transactionRequest.Amount
	repository.UpdateAccount(originAccount)
	repository.UpdateAccount(destinyAccount)

	transaction := &entities.Transaction{
		OriginAccountNumber:  transactionRequest.OriginAccountNumber,
		DestinyAccountNumber: transactionRequest.DestinyAccountNumber,
		Amount:               transactionRequest.Amount,
		AccountID:            claims.Id,
	}

	errTx := repository.CreateTransaction(transaction)
	if errTx != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating transaction",
		})
		return
		//TODO: Implement balance rollback
	}

	ctx.IndentedJSON(http.StatusAccepted, gin.H{
		"message": "Transaction performed successfully",
	})

}

func TokenHandleTransaction(ctx *gin.Context) *authApplication.Claims {

	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "empty auth header",
		})
		return nil
	}

	valid, errToken := authApplication.ValidateToken(token)
	if errToken != nil || !valid {
		ctx.IndentedJSON(http.StatusForbidden, gin.H{
			"message": "Not allowed token",
		})
		return nil
	}

	claims, errDecode := authApplication.DecodeToken(token)
	if errDecode != nil {
		ctx.IndentedJSON(http.StatusForbidden, gin.H{
			"message": "Not found token",
		})
		return nil
	}

	return claims
}

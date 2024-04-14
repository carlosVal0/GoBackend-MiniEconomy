package infraestructure

import (
	"net/http"

	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/application"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	authApplication "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/application"
	"github.com/gin-gonic/gin"
)

type AccountDto struct {
	AccountNumber  string  `json:"accountNumber"`
	CurrentBalance float64 `json:"balance"`
}

type RechargeDto struct {
	AccountNumber string  `json:"accountNumber"`
	Amount        float64 `json:"amount"`
}

func CreateAccount(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "empty auth header",
		})
		return
	}

	valid, errToken := authApplication.ValidateToken(token)
	if errToken != nil || !valid {
		ctx.IndentedJSON(http.StatusForbidden, gin.H{
			"message": "Not allowed token",
		})
		return
	}

	claims, errDecode := authApplication.DecodeToken(token)
	if errDecode != nil {
		ctx.IndentedJSON(http.StatusForbidden, gin.H{
			"message": "Not found token",
		})
		return
	}

	account, errCreate := application.CreateAccountService(claims.Id)
	if errCreate != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Error creating account",
		})
		return
	}

	accDto := &AccountDto{
		AccountNumber:  account.AccountNumber,
		CurrentBalance: account.CurrentBalance,
	}

	ctx.IndentedJSON(http.StatusCreated, accDto)

}

func GetAccounts(ctx *gin.Context) {

	claims := TokenHandle(ctx)
	if claims == nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "empty auth header",
		})
		return
	}

	accs, err := application.GetAccountsService(claims.Id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "empty auth header",
		})
		return
	}

	accsDto := MapAccountDto(accs)

	ctx.IndentedJSON(http.StatusOK, accsDto)

}

func RechargeAccount(ctx *gin.Context) {
	claims := TokenHandle(ctx)
	if claims == nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "empty auth header",
		})
		return
	}
	var recharge RechargeDto
	err := ctx.BindJSON(&recharge)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Wrong request",
		})
		return
	}

	recError := application.RechargeAccount(claims.Id, recharge.AccountNumber, recharge.Amount)
	if recError != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error recharging account",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{
		"message": "Succesfull recharge",
	})

}

func TokenHandle(ctx *gin.Context) *authApplication.Claims {

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

func MapAccountDto(accs []entities.Account) []AccountDto {
	accountsDto := make([]AccountDto, 0)

	for _, account := range accs {
		tempAccDto := &AccountDto{
			AccountNumber:  account.AccountNumber,
			CurrentBalance: account.CurrentBalance,
		}

		accountsDto = append(accountsDto, *tempAccDto)
	}

	return accountsDto
}

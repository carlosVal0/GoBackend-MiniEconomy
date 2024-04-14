package main

import (
	"log"
	"os"

	accountInfraestructure "github.com/carlosVal0/miniEconomyGoBackend/modules/account/infraestructure"
	authRepository "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/repository"
	authInfraestructure "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/infraestructure"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()
	logger := log.Default()
	logger.SetPrefix("main-minieconomy at ")
	authRepository.Connect()
	os.Setenv("TZ", "America/Bogota")
}

func main() {

	Init()

	router := gin.Default()

	router.POST("/login", authInfraestructure.LoginController)
	router.POST("/register", authInfraestructure.RegisterController)
	router.POST("/decode", authInfraestructure.DecodingController)
	router.POST("/account", accountInfraestructure.CreateAccount)
	router.GET("/accounts", accountInfraestructure.GetAccounts)
	router.POST("/recharge", accountInfraestructure.RechargeAccount)

	router.Run("localhost:8080")

}

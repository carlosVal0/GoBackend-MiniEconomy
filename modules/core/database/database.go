package database

import (
	"fmt"
	"log"
	"os"

	accountDomain "github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	authDomain "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbBean *gorm.DB

func Connect() {

	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("%v:%v@tcp(%v)/%v?parseTime=true", dbUser, dbPasswd, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dbBean = db
	dbBean.AutoMigrate(&authDomain.User{}, &accountDomain.Account{}, &accountDomain.Transaction{})
	logger := log.Default()
	logger.Println("Connected to " + dbBean.Name())

}

func GetDbBean() *gorm.DB {
	return dbBean
}

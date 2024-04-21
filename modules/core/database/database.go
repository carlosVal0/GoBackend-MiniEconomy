package database

import (
	"fmt"
	"log"
	"os"

	accountDomain "github.com/carlosVal0/miniEconomyGoBackend/modules/account/domain/entities"
	authDomain "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbBean *gorm.DB

func Connect() {

	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSchema := os.Getenv("DB_SCHEMA")
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v port=%v host=%v TimeZone=America/Bogota", dbUser, dbPasswd, dbName, dbPort, dbHost)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   dbSchema,
			SingularTable: false,
		},
	})

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

package repository

import (
	"fmt"
	"log"
	"os"
	"time"

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

func CreateUser(u *authDomain.User) error {

	result := dbBean.Select("name", "email", "password", "orgId").Create(&u)
	return result.Error

}

func GetUser(id int) (*authDomain.User, error) {
	var user authDomain.User
	result := dbBean.First(&user, id).Omit("password")
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func DeleteUser(id int) error {
	result := dbBean.Delete(&authDomain.User{}, id)
	return result.Error
}

func UpdateUser(u *authDomain.User) (*authDomain.User, error) {
	dbBean.Model(&u).Updates(&authDomain.User{
		ID:        u.ID,
		Name:      u.Name,
		Password:  u.Password,
		UpdatedAt: time.Now(),
		CreatedAt: u.CreatedAt,
		Email:     u.Email,
		OrgId:     u.OrgId,
	})
	result := dbBean.Save(&u)
	return u, result.Error

}

func GetUserByEmailPasswd(email string, password string) (*authDomain.User, error) {

	var user *authDomain.User
	result := dbBean.Where(" email = ? AND password = ?", email, password).First(&user)

	return user, result.Error

}

package repository

import (
	"time"

	authDomain "github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/entities"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/core/database"
)

func CreateUser(u *authDomain.User) error {

	dbBean := database.GetDbBean()

	result := dbBean.Select("name", "email", "password", "orgId").Create(&u)
	return result.Error

}

func GetUser(id int) (*authDomain.User, error) {

	dbBean := database.GetDbBean()

	var user authDomain.User
	result := dbBean.First(&user, id).Omit("password")
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil

}

func DeleteUser(id int) error {

	dbBean := database.GetDbBean()

	result := dbBean.Delete(&authDomain.User{}, id)
	return result.Error
}

func UpdateUser(u *authDomain.User) (*authDomain.User, error) {

	dbBean := database.GetDbBean()

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

	dbBean := database.GetDbBean()

	var user *authDomain.User
	result := dbBean.Where(" email = ? AND password = ?", email, password).First(&user)

	return user, result.Error

}

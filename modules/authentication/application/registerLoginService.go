package application

import (
	"errors"
	"os"
	"time"

	"github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/entities"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/repository"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

type LoggedUser struct {
	ExpiresAt int    `json:"expiresAt"`
	Token     string `json:"token"`
}

func RegisterUser(u *entities.User) error {

	if u.Password == "" || u.Email == "" || u.Name == "" {
		return errors.New("not allowed password, email or name go null")
	}

	err := repository.CreateUser(u)
	return err
}

func LoginUser(email string, password string) (*LoggedUser, error) {

	godotenv.Load()

	if email == "" || password == "" {
		return nil, errors.New("not allowed null passwords or emails")
	}

	user, err := repository.GetUserByEmailPasswd(email, password)
	if err != nil {
		return nil, err
	}

	expiringTime := time.Now().Add(time.Hour * time.Duration(1))

	claims := Claims{
		Id:    user.ID,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiringTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, err
	}

	loggedUser := &LoggedUser{
		ExpiresAt: int(time.Hour) / int(time.Second),
		Token:     tokenStr,
	}

	return loggedUser, nil

}

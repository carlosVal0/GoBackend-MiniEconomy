package infraestructure

import (
	"net/http"
	"time"

	"github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/application"
	"github.com/carlosVal0/miniEconomyGoBackend/modules/authentication/domain/entities"
	"github.com/gin-gonic/gin"
)

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginController(ctx *gin.Context) {
	var loginDto LoginDto

	if err := ctx.BindJSON(&loginDto); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error interpreting request",
		})
		return
	}

	loggedUser, err := application.LoginUser(loginDto.Email, loginDto.Password)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error interpreting request",
		})
		return
	}

	ctx.IndentedJSON(http.StatusAccepted, loggedUser)
}

type RegisterUserDto struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	OrgId    int    `json:"orgId,omitempty"`
}

func DecodingController(ctx *gin.Context) {

	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error interpreting request",
		})
		return
	}

	claims, err := application.DecodeToken(token)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error interpreting request",
		})
		return
	}

	ctx.IndentedJSON(http.StatusOK, claims)

}

func RegisterController(ctx *gin.Context) {

	var registerUser RegisterUserDto

	if err := ctx.BindJSON(&registerUser); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error interpreting request",
		})
		return
	}

	user := &entities.User{
		Name:      registerUser.Name,
		Email:     registerUser.Email,
		Password:  registerUser.Password,
		OrgId:     registerUser.OrgId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := application.RegisterUser(user)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Error interpreting request",
		})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Successfull register, check your email",
	})
}

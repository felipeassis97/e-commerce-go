package controllers

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	usecase usecase.LoginUseCase
}

func NewLoginController(usecase usecase.LoginUseCase) LoginController {
	return LoginController{
		usecase: usecase,
	}
}

func (u *LoginController) Login(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response := model.Response{
			Type:    model.InvalidBody,
			Message: "Invalid body credentials",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	login, err := u.usecase.Login(user)
	if err != nil {
		response := model.Response{
			Type:    model.Unauthorized,
			Message: "Login error. Check your credentials",
		}
		ctx.JSON(http.StatusUnauthorized, response)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User found",
		"user":    login,
	})
}

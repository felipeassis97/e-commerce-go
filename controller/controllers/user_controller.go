package controllers

import (
	"fmt"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecase.UserUseCase
}

func NewUserController(usecase usecase.UserUseCase) UserController {
	return UserController{
		userUseCase: usecase,
	}
}

func (u *UserController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("userID")
	if id == "" {
		response := model.Response{
			Type:    model.MissingParams,
			Message: "Invalid user ID",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	userID, err := strconv.Atoi(id)
	if err != nil {
		response := model.Response{
			Type:    model.MissingParams,
			Message: "Invalid user ID",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	users, err := u.userUseCase.GetUserByID(userID)

	if err != nil {
		response := model.Response{
			Type:    model.InternalError,
			Message: fmt.Sprintf("Error getting user ID: %s", id),
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user model.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		response := model.Response{
			Type:    model.InvalidBody,
			Message: "Error parsing request body",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	createdUser, err := u.userUseCase.CreateUser(user)
	if err != nil {
		response := model.Response{
			Type:    model.InternalError,
			Message: "Error while creating product",
		}
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"user":    createdUser,
	})
}

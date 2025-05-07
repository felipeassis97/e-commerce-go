package routes

import (
	"go-api/controller/controllers"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(server *gin.Engine, controller controllers.UserController) {
	users := server.Group("/user")
	{
		users.GET("/:userID", controller.GetUserByID)
		users.POST("/create", controller.CreateUser)
	}
}

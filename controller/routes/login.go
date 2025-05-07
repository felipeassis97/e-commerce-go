package routes

import (
	"go-api/controller/controllers"

	"github.com/gin-gonic/gin"
)

func registerLoginRoutes(server *gin.Engine, controller controllers.LoginController) {
	server.POST("/login", controller.Login)
}

package routes

import (
	"database/sql"
	"go-api/controller/controllers"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes register all application routes
func RegisterRoutes(server *gin.Engine, conn *sql.DB) {
	registerUserRoutes(server, buildUserController(conn))
	registerProductRoutes(server, buildProductController(conn))
}

func buildProductController(conn *sql.DB) controllers.ProductController {
	repo := repository.NewProductRepository(conn)
	useCase := usecase.NewProductUsecase(repo)
	return controllers.NewProductController(useCase)
}

func buildUserController(conn *sql.DB) controllers.UserController {
	repo := repository.NewUserRepository(conn)
	useCase := usecase.NewUserUseCase(repo)
	return controllers.NewUserController(useCase)
}

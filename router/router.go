package router

import (
	"blogging-platform/controller"
	"blogging-platform/repository"
	"blogging-platform/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(DB *gorm.DB) *gin.Engine {
	// Setup repository, service, dan controller
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	app := gin.Default()

	// Routing
	userRouter := app.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	return app
}

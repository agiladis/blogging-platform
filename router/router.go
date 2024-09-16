package router

import (
	"blogging-platform/controller"
	"blogging-platform/middleware"
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

	blogPostRepository := repository.NewBlogPostRepository(DB)
	blogPostService := service.NewBlogPostService(blogPostRepository)
	blogPostController := controller.NewBlogPostController(blogPostService)

	app := gin.Default()

	// Routing
	userRouter := app.Group("/users")
	{
		userRouter.POST("/register", userController.Register)
		userRouter.POST("/login", userController.Login)
	}

	blogPostRouter := app.Group("/posts")
	{
		blogPostRouter.Use(middleware.JWTMiddleware())
		blogPostRouter.POST("/", blogPostController.CreateBlogPost)
		blogPostRouter.GET("/", blogPostController.GetAllPosts)
		blogPostRouter.GET("/:id", blogPostController.GetPostById)
	}

	return app
}

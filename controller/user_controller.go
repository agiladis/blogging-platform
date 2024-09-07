package controller

import (
	"blogging-platform/dto"
	"blogging-platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{userService}
}

func (c *userController) Register(ctx *gin.Context) {
	var userRegisterDTO dto.UserRegisterDTO

	if err := ctx.ShouldBindJSON(&userRegisterDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	user, err := c.userService.Register(userRegisterDTO)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"message": "user created",
		"data": user,
	})
}
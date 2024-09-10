package controller

import (
	"blogging-platform/dto"
	"blogging-platform/helper"
	"blogging-platform/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type blogPostController struct {
	blogPostService service.BlogPostService
}

func NewBlogPostController(blogPostService service.BlogPostService) *blogPostController {
	return &blogPostController{blogPostService}
}

func (c *blogPostController) CreateBlogPost(ctx *gin.Context) {
	var blogPostDTO dto.BlogPostDTO

	if err := ctx.ShouldBindBodyWithJSON(&blogPostDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
	}

	// get user id from ctx
	accessClaim, err := helper.GetIdentityFromCtx(ctx)
	if err != nil {
		return
	}

	// hit service
	blogPost, err := c.blogPostService.Create(blogPostDTO, uint(accessClaim.AccessClaims.ID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "user created",
		"data":    blogPost,
	})
}

package controller

import (
	"blogging-platform/dto"
	"blogging-platform/helper"
	"blogging-platform/service"
	"net/http"
	"strconv"

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

	if err := ctx.ShouldBindJSON(&blogPostDTO); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
			"data":    nil,
		})
		return
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
		"message": "blog post created",
		"data":    blogPost,
	})
}

func (c *blogPostController) GetAllPosts(ctx *gin.Context) {
	blogPosts, err := c.blogPostService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "success",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get all post success",
		"data":    blogPosts,
	})
}

func (c *blogPostController) GetPostById(ctx *gin.Context) {
	blogPostId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "invalid ID",
			"data":    nil,
		})
		return
	}

	blogPost, err := c.blogPostService.GetById(uint(blogPostId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "post not found",
			"data":    nil,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "get post success",
		"data":    blogPost,
	})
}

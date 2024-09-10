package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIdentityFromCtx(ctx *gin.Context) (dataOut Claims, err error) {
	accessClaim, ok := ctx.Get("access_claim")
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "please logn first",
		})
		return
	}

	if err = ObjectMapper(accessClaim, &dataOut); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": "invalid token payload",
		})
		return
	}

	return
}

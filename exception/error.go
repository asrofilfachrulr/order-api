package exception

import (
	"github.com/gin-gonic/gin"
)

func BadRequestError(c *gin.Context, err error) {
	if err == nil {
		return
	}
	c.AbortWithStatusJSON(400, gin.H{
		"status":  "Bad Request",
		"message": err.Error(),
	})
	panic(err)
	// c.JSON(400, gin.H{
	// 	"status":  "Bad Request",
	// 	"message": err.Error(),
	// })
	// <-c.Request.Context().Done()
}

func NotFoundError(c *gin.Context, err error) {
	c.JSON(404, gin.H{
		"status":  "Not found",
		"message": err.Error(),
	})
}

func InternalServerError(c *gin.Context, err error) {
	c.JSON(500, gin.H{
		"status":  "Internal Error",
		"message": err.Error(),
	})
}

package exception

import (
	"log"

	"github.com/gin-gonic/gin"
)

func TerminateRuntimeIfError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func BadRequestError(c *gin.Context, ers ...error) {
	var msg error
	if ers[0] == nil {
		return
	}

	if len(ers) > 1 {
		msg = ers[1]
	} else {
		msg = ers[0]
	}
	c.AbortWithStatusJSON(400, gin.H{
		"status":  "Error: Bad Request",
		"message": msg.Error(),
	})
	panic(msg)
}

func NotFoundError(c *gin.Context, ers ...error) {
	var msg error
	if ers[0] == nil {
		return
	}

	if len(ers) > 1 {
		msg = ers[1]
	} else {
		msg = ers[0]
	}
	c.JSON(404, gin.H{
		"status":  "Error: Not found",
		"message": msg.Error(),
	})
	panic(msg)
}

func InternalServerError(c *gin.Context, ers ...error) {
	var msg error
	if ers[0] == nil {
		return
	}

	if len(ers) > 1 {
		msg = ers[1]
	} else {
		msg = ers[0]
	}
	c.JSON(500, gin.H{
		"status":  "Error: Internal Error",
		"message": msg.Error(),
	})
	panic(msg)
}

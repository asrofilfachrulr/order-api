package util

import (
	"log"
	"orderapi/model"

	"github.com/gin-gonic/gin"
)

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func ExitIfError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func BadInputErrorResp(e error, c *gin.Context) {
	if e != nil {
		c.JSON(405, model.Response{
			Status:  "Bad Input",
			Message: e.Error(),
		})
	}
}

func InternalErrorResp(e error, c *gin.Context) {
	if e != nil {
		c.JSON(500, model.Response{
			Status:  "Internal Error",
			Message: e.Error(),
		})
	}
}

package util

import (
	"orderapi/model"

	"github.com/gin-gonic/gin"
)

func PanicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func BadInputErrorResp(e error, c *gin.Context) {

}

func InternalErrorResp(e error, c *gin.Context) {
	if e != nil {
		c.JSON(500, model.Response{
			Status:  "Internal Error",
			Message: e.Error(),
		})
	}
}

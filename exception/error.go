package exception

import (
	"log"
	berror "orderapi/error" // built in error

	"github.com/gin-gonic/gin"
)

func CheckCaseErrorThenRespond(c *gin.Context, e berror.Error) {
	switch e.(type) {
	case *berror.BadRequestError:
		RespondWithBadRequestError(c, e.GetError())
	case *berror.NotFoundError:
		RespondWithNotFoundError(c, e.GetError())
	case *berror.InternalServerError:
		RespondWithInternalServerError(c, e.GetError())
	default:
		panic(e) // terminate current handler flow
	}

}

func TerminateRuntimeIfError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func RespondWithBadRequestError(c *gin.Context, e error) {
	c.AbortWithStatusJSON(400, gin.H{
		"status":  "Error: Bad Request",
		"message": e.Error(),
	})
	panic(e) // terminate current handler flow
}

func RespondWithNotFoundError(c *gin.Context, e error) {

	c.JSON(404, gin.H{
		"status":  "Error: Not found",
		"message": e.Error(),
	})
	panic(e) // terminate current handler flow
}

func RespondWithInternalServerError(c *gin.Context, e error) {

	c.JSON(500, gin.H{
		"status":  "Error: Internal Error",
		"message": e.Error(),
	})
	panic(e) // terminate current handler flow
}

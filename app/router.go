package app

import (
	"orderapi/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Home",
		})
	})

	v1 := r.Group("/api/v1")
	{
		order := v1.Group("/order")
		{
			order.GET("/:orderId", h.GetOrderById)
			order.PUT("/:orderId", h.PutOrderById)
			order.POST("/", h.PostOrder)
			order.DELETE("/:orderId", h.DeleteOrderById)
		}

		orders := v1.Group("/orders")
		{
			orders.GET("/", h.GetMultipleOrder)
			orders.PUT("/", h.PutMultipleOrder)
			orders.DELETE("/", h.DeleteMultipleOrder)
		}
	}

	return r
}

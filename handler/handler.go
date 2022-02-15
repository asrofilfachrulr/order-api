package handler

import "github.com/gin-gonic/gin"

type HandleFunc struct{}

func NewHandler() *HandleFunc {
	return &HandleFunc{}
}

// Order
func (h *HandleFunc) PostOrder(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	// Parse JSON from request body!
	c.JSON(200, gin.H{
		"message": "You're at PostOrder",
	})
}
func (h *HandleFunc) GetOrderById(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at GetOrderById by id =" + c.Param("orderId"),
	})
}
func (h *HandleFunc) PutOrderById(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at PutOrderById by id =" + c.Param("orderId"),
	})
}
func (h *HandleFunc) DeleteOrderById(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at DeleteOrderById by id =" + c.Param("orderId"),
	})
}

// Orders
func (h *HandleFunc) GetMultipleOrder(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at GetMultipleOrder",
	})
}
func (h *HandleFunc) DeleteMultipleOrder(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at DeleteMultipleOrder",
	})
}
func (h *HandleFunc) PutMultipleOrder(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at PutMultipleOrder",
	})
}

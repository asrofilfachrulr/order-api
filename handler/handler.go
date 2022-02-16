package handler

import (
	"orderapi/controller"
	"orderapi/model"
	"orderapi/util"

	"github.com/gin-gonic/gin"
)

type HandleFunc struct{}

func NewHandler() *HandleFunc {
	return &HandleFunc{}
}

// Order
func (h *HandleFunc) PostOrder(c *gin.Context) {
	jsonData, err := (&model.Order{}).ParseJSON(c.Request.Body)
	util.PanicIfError(err)

	resp, err := controller.MakeOrder(jsonData)
	util.InternalErrorResp(err, c)

	c.JSON(200, *resp)
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

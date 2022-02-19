package handler

import (
	"orderapi/controller"
	"orderapi/model"
	"orderapi/util"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Controller *controller.Controller
}

func NewHandler(c *controller.Controller) *Handler {
	return &Handler{
		Controller: c,
	}
}

// Order
func (h *Handler) PostOrder(c *gin.Context) {
	jsonData, err := (&model.Order{}).ParseJSON(c.Request.Body)
	util.BadInputErrorResp(err, c)

	resp, err := h.Controller.MakeOrder(&jsonData)
	util.BadInputErrorResp(err, c)

	c.JSON(200, resp)
}
func (h *Handler) GetOrderById(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at GetOrderById by id =" + c.Param("orderId"),
	})
}
func (h *Handler) PutOrderById(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at PutOrderById by id =" + c.Param("orderId"),
	})
}
func (h *Handler) DeleteOrderById(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at DeleteOrderById by id =" + c.Param("orderId"),
	})
}

// Orders
func (h *Handler) GetMultipleOrder(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at GetMultipleOrder",
	})
}
func (h *Handler) DeleteMultipleOrder(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at DeleteMultipleOrder",
	})
}
func (h *Handler) PutMultipleOrder(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	c.JSON(200, gin.H{
		"message": "You're at PutMultipleOrder",
	})
}

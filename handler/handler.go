package handler

import (
	"log"
	"orderapi/controller"
	"orderapi/exception"
	"orderapi/helper"
	"orderapi/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	Controller *controller.Controller
	Validate   *validator.Validate
}

func NewHandler(c *controller.Controller, v *validator.Validate) *Handler {
	return &Handler{
		Controller: c,
		Validate:   v,
	}
}

func EmptyRecover() {
	if r := recover(); r != nil {
		log.Printf("recovered from error [%s]", r)
	}
}

// Order
func (h *Handler) PostOrder(c *gin.Context) {
	defer EmptyRecover()
	OrderStruct, err := (&model.Order{}).ParseJSON(c.Request.Body)
	exception.BadRequestError(c, err)

	log.Println(OrderStruct)

	err = h.Validate.Struct(OrderStruct)
	exception.BadRequestError(c, err)

	h.Controller.MakeOrder(c, &OrderStruct)

	c.JSON(201, gin.H{
		"status":  "success",
		"message": "succeed register new order",
		"data":    helper.ToOrderResponse(&OrderStruct),
	})
}
func (h *Handler) GetOrderById(c *gin.Context) {
	// TODO: implement this are you kidding me for 2000 bucks
	orderStruct := &model.Order{}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success retrieve order " + c.Param("orderId"),
		"data":    *orderStruct,
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

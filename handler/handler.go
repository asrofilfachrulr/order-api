package handler

import (
	"errors"
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

func EmptyRecover(c *gin.Context) {
	if r := recover(); r != nil {
		log.Printf("recovered from error [%s]", r)
		exception.RespondWithInternalServerError(c, errors.New(r.(string)))
	}
}

// Order
func (h *Handler) PostOrder(c *gin.Context) {
	defer EmptyRecover(c)
	OrderStruct, err := (&model.Order{}).ParseJSON(c.Request.Body)
	if err != nil {
		exception.RespondWithBadRequestError(c, err)
	}

	log.Println(OrderStruct)

	err = h.Validate.Struct(OrderStruct)
	if err != nil {
		exception.RespondWithBadRequestError(c, errors.New("wrong attribute type"))
	}

	e := h.Controller.MakeOrder(&OrderStruct)
	if e != nil {
		exception.CheckCaseErrorThenRespond(c, e)
	}

	c.JSON(201, gin.H{
		"status":  "success",
		"message": "succeed register new order",
		"data":    helper.ToOrderResponse(&OrderStruct),
	})
}
func (h *Handler) GetOrderById(c *gin.Context) {
	defer EmptyRecover(c)
	orderDetail, e := h.Controller.GetOrderById(c.Param("orderId"))
	if e != nil {
		exception.CheckCaseErrorThenRespond(c, e)
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success retrieve order " + c.Param("orderId"),
		"data":    *orderDetail,
	})
}
func (h *Handler) PutOrderById(c *gin.Context) {
	log.Println("Entering PutOrderById block")
	defer EmptyRecover(c)
	j, err := (&model.OrderUpdate{}).ParseJSON(c.Request.Body)
	if err != nil {
		exception.RespondWithBadRequestError(c, err)
	}

	h.Validate.RegisterValidation("status", model.ValidateStatus)
	err = h.Validate.Struct(j)
	if err != nil {
		exception.RespondWithBadRequestError(c, err)
	}

	err = j.ValidateMissing()
	if err != nil {
		exception.RespondWithBadRequestError(c, err)
	}

	log.Println("Ending PutOrderById block")
	c.JSON(204, gin.H{
		"status":  "success",
		"message": "Success updating order " + c.Param("orderId"),
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

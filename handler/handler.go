package handler

import (
	"errors"
	"log"
	"orderapi/controller"
	"orderapi/exception"
	"orderapi/helper"
	"orderapi/model"
	"orderapi/utils"

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
		"data":    helper.ToOrderCreatedResponse(&OrderStruct),
	})
}
func (h *Handler) GetOrderById(c *gin.Context) {
	defer EmptyRecover()
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
	defer EmptyRecover()
	id := c.Param("orderId")

	e := h.Controller.CheckOrderId(id)
	if e != nil {
		exception.CheckCaseErrorThenRespond(c, e)
	}

	j, err := (&model.OrderUpdate{}).ParseJSON(c.Request.Body)
	if err != nil {
		exception.RespondWithBadRequestError(c, err)
	}

	err = j.CheckParsedJSON()
	if err != nil {
		exception.RespondWithBadRequestError(c, err)
	}

	h.Validate.RegisterValidation("status-validate", model.ValidateStatus)

	err = h.Validate.Struct(j)
	if err != nil {
		exception.RespondWithBadRequestError(c, err)
	}

	log.Println(j)

	e = h.Controller.UpdateOrderById(id, &j)
	if e != nil {
		exception.CheckCaseErrorThenRespond(c, e)
	}

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success updating order " + id,
	})
}
func (h *Handler) DeleteOrderById(c *gin.Context) {
	defer EmptyRecover()

	id := c.Param("orderId")

	e := h.Controller.CheckOrderId(id)
	if e != nil {
		exception.CheckCaseErrorThenRespond(c, e)
	}
	e = h.Controller.DeleteOrderById(id)
	if e != nil {
		exception.CheckCaseErrorThenRespond(c, e)
	}
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "Success deleting id =" + id,
	})
}

// Orders
func (h *Handler) GetMultipleOrder(c *gin.Context) {
	defer EmptyRecover()
	f := &model.Filter{}

	// getting query params value
	f.Status, _ = c.GetQuery("status")
	f.From, _ = c.GetQuery("from")
	f.To, _ = c.GetQuery("to")

	f.ValidateStatus()
	f.ValidateTime()

	orders, e := h.Controller.GetAllOrder(f)
	if e != nil {
		exception.CheckCaseErrorThenRespond(c, e)
	}
	data := helper.ToOrdersBriefResponse(orders)
	count := utils.IntToString(len(data))
	c.JSON(200, gin.H{
		"status":  "success",
		"message": "success retrieving all data: " + count + " item(s)",
		"data":    data,
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

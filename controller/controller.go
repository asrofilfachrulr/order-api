package controller

import (
	"errors"
	"orderapi/exception"
	"orderapi/inmemory"
	"orderapi/model"
	"orderapi/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	nanoid "github.com/matoous/go-nanoid/v2"
)

type Controller struct {
	Service *service.OrderService
}

func NewController(s *service.OrderService) *Controller {
	return &Controller{
		Service: s,
	}
}

// /order
func (c *Controller) MakeOrder(ctx *gin.Context, o *model.Order) {
	// order items cannot be empty
	if len(o.Items) == 0 {
		exception.BadRequestError(ctx, errors.New("cannot register new order due order items is zero"))
	}

	// generate unique id
	id, _ := nanoid.New()
	o.Id = id

	// add (or maybe reset for safety) timestamp for every make order request
	o.CreatedAt = time.Now()
	// ensure that status must be unpaid
	o.Status = "unpaid"
	// calculate manually the total
	total := 0
	for _, item := range o.Items {
		if item.Qty < 1 {
			exception.BadRequestError(ctx, errors.New("item's quantity must be at least 1"))
		}
		if _, f := inmemory.ListMenuInmemory[item.MenuId]; !f {
			exception.BadRequestError(ctx, errors.New("Menu "+strconv.Itoa(item.MenuId)+" not found"))
			return

		}
		total += inmemory.ListMenuInmemory[item.MenuId].Price * int(item.Qty)
	}
	o.Total = int64(total)

	c.Service.AddOrder(ctx, o)
}

func (c *Controller) GetOrderById(ctx *gin.Context, o *model.Order) {

}

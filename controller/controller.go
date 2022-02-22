package controller

import (
	"errors"
	"orderapi/error"
	"orderapi/inmemory"
	"orderapi/model"
	"orderapi/service"
	"strconv"
	"time"

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
func (c *Controller) MakeOrder(o *model.Order) error.Error {
	// order items cannot be empty
	if len(o.Items) == 0 {
		return &error.BadRequestError{Err: errors.New("cannot register new order due order items is zero")}
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
			return &error.BadRequestError{Err: errors.New("item's quantity must be at least 1")}
		}
		if _, f := inmemory.ListMenuInmemory[item.MenuId]; !f {
			return &error.BadRequestError{Err: errors.New("Menu " + strconv.Itoa(item.MenuId) + " not found")}

		}
		total += inmemory.ListMenuInmemory[item.MenuId].Price * int(item.Qty)
	}
	o.Total = int64(total)

	c.Service.AddOrder(o)
	return nil
}

func (c *Controller) GetOrderById(o *model.Order) error.Error {

	return nil
}

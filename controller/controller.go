package controller

import (
	"errors"
	"orderapi/inmemory"
	"orderapi/model"
	"orderapi/service"
	"strconv"
	"time"
)

type Controller struct {
	Service *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{
		Service: s,
	}
}

// /order
func (c *Controller) MakeOrder(o *model.Order) (*model.Response, error) {
	// add (or maybe reset for safety) timestamp for every make order request
	o.CreatedAt = time.Now()
	// ensure that status must be unpaid
	o.Status = "unpaid"
	// calculate manually the total
	total := 0
	for _, item := range o.Items {
		if _, f := inmemory.ListMenuInmemory[item.MenuId]; !f {
			return nil, errors.New("menuId " + strconv.Itoa(item.MenuId) + " not found")
		}
		total += inmemory.ListMenuInmemory[item.MenuId] * int(item.Qty)
	}
	o.Total = int64(total)

	return &model.Response{
		Status:  "success",
		Message: "success add new order",
		Data:    *o,
	}, nil
}

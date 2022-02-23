package controller

import (
	"errors"
	"orderapi/error"
	"orderapi/inmemory"
	"orderapi/model"
	"orderapi/service"
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
	o.UpdatedAt = o.CreatedAt
	// ensure that status must be unpaid
	o.Status = "unpaid"
	// calculate manually the total
	total := 0
	for _, item := range o.Items {
		if item.Qty < 1 {
			return &error.BadRequestError{Err: errors.New("item's quantity must be at least 1")}
		}
		if _, f := inmemory.ListMenuInmemory[item.Name]; !f {
			return &error.BadRequestError{Err: errors.New("Menu " + item.Name + " not found")}

		}
		total += inmemory.ListMenuInmemory[item.Name].Price * int(item.Qty)
	}
	o.Total = int64(total)

	err := c.Service.AddOrder(o)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetOrderById(id string) (*model.OrderDetailResp, error.Error) {
	var order *model.Order

	order, err := c.Service.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	orderDetail := &model.OrderDetailResp{
		Id:        order.Id,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
		Status:    order.Status,
		Total:     order.Total,
	}

	var items []model.OrderItemDetailResp
	for _, item := range order.Items {
		i := model.OrderItemDetailResp{
			Qty:  item.Qty,
			Name: item.Name,
		}
		items = append(items, i)
	}

	orderDetail.Items = items

	return orderDetail, nil
}

func (c *Controller) UpdateOrderStatusById(id string) error.Error {
	err := c.Service.UpdateOrderStatusById(id)
	if err != nil {
		return err
	}
	return nil
}

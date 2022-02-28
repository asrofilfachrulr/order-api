package controller

import (
	"errors"
	"log"
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
	o.UpdatedAt = o.CreatedAt
	// ensure that status must be unpaid
	o.Status = "unpaid"
	// calculate manually the total
	total := 0
	for _, item := range o.Items {
		if item.Qty < 1 {
			return &error.BadRequestError{Err: errors.New("item's quantity must be at least 1")}
		}
		if _, f := inmemory.ListMenuInmemory[item.Id]; !f {
			return &error.BadRequestError{Err: errors.New("Menu " + strconv.Itoa(item.Id) + " not found")}

		}
		total += inmemory.ListMenuInmemory[item.Id].Price * int(item.Qty)
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
			Name: inmemory.ListMenuInmemory[item.Id].Name,
		}
		items = append(items, i)
	}

	orderDetail.Items = items

	return orderDetail, nil
}

func (c *Controller) UpdateOrderById(id string, o *model.OrderUpdate) error.Error {
	statusUpdated := false
	itemsUpdated := false

	// check for status attribute existence then execute change if exists
	if o.Status != "" {
		log.Println("Updating status")
		err := c.Service.UpdateOrderStatusById(id, o.Status)
		if err != nil {
			return err
		}
		log.Println("Status updated")
		statusUpdated = true
	}
	// check for items attribute existence then execute change if exists
	if len(o.Items) != 0 {
		log.Println("Updating items")
		err := c.Service.UpdateOrderItemById(id, o.Items)
		if err != nil {
			return err
		}
		log.Println("Items updated")
		itemsUpdated = true
	}

	if statusUpdated || itemsUpdated {
		log.Println("Updating stamp")
		err := c.Service.UpdateStamp(id)
		if err != nil {
			return err
		}
		log.Println("Stamp updated")
		if itemsUpdated {
			log.Println("Updating total")
			err := c.Service.UpdateTotal(id)
			if err != nil {
				return err
			}
			log.Println("Total updated")
		}
	}

	return nil
}

func (c *Controller) CheckOrderId(id string) error.Error {
	err := c.Service.CheckOrderId(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) DeleteOrderById(id string) error.Error {
	err := c.Service.DeleteOrderById(id)
	if err != nil {
		return err
	}
	return nil
}

func (c *Controller) GetAllOrder(f *model.Filter) ([]model.Order, error.Error) {
	orders := []model.Order{}
	err := c.Service.GetAllOrder(f, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *Controller) UpdateOrdersStatusByDate(f *model.Filter) error.Error {
	err := c.Service.UpdateOrdersStatusByDate(f)
	if err != nil {
		return err
	}
	return nil
}
func (c *Controller) DeleteOrdersByFilter(f *model.Filter) error.Error {
	err := c.Service.DeleteOrdersByFilter(f)
	if err != nil {
		return err
	}
	return nil
}

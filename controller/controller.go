package controller

import (
	"errors"
	"log"
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
func (c *Controller) MakeOrder(o *model.Order) (gin.H, error) {
	// generate unique id
	id, _ := nanoid.New()
	o.Id = "order-" + id

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

	log.Println(o)

	err := c.Service.AddOrder(o)

	if err != nil {
		return nil, err
	}

	return gin.H{
		"status":  "success",
		"message": "succeed register your order",
		"id":      o.Id,
	}, nil
}

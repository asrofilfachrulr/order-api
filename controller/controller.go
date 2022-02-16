package controller

import (
	"errors"
	"orderapi/inmemory"
	"orderapi/model"
	"time"
)

// /order
func MakeOrder(o *model.Order) (*model.Response, error) {
	// add (or maybe reset for safety) timestamp for every make order request
	o.CreatedAt = time.Now()
	// ensure that status must be unpaid
	o.Status = "unpaid"
	// calculate manually the total
	total := 0
	for _, item := range o.Items {
		if _, f := inmemory.ListMenuRuntime[item.MenuId]; !f {
			return nil, errors.New("menuId " + item.MenuId + " not found")
		}
		total += inmemory.ListMenuRuntime[item.MenuId] * int(item.Qty)
	}
	o.Total = int64(total)

	return &model.Response{
		Status:  "success",
		Message: "success add new order",
		Data:    *o,
	}, nil
}

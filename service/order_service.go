package service

import (
	"database/sql"
	"orderapi/error"
	"orderapi/model"
)

type OrderService struct {
	DB *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		DB: db,
	}
}

func (o *OrderService) AddOrder(order *model.Order) error.Error {
	q := "INSERT INTO order_list VALUES($1, $2, $3, $4)"
	_, err := o.DB.Exec(q, order.Id, order.CreatedAt, order.Status, order.Total)

	if err != nil {
		return &error.InternalServerError{Err: err}
	}

	for _, item := range order.Items {
		q = "INSERT INTO order_item VALUES($1, $2, $3)"
		_, err := o.DB.Exec(q, order.Id, item.MenuId, item.Qty)

		if err != nil {
			return &error.InternalServerError{Err: err}
		}
	}
	return nil
}

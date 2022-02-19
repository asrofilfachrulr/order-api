package service

import (
	"database/sql"
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

func (o *OrderService) AddOrder(order *model.Order) error {
	q := "INSERT INTO order_list VALUES(?, ?, ?, ?)"
	_, err := o.DB.Exec(q, order.Id, order.CreatedAt, order.Status, order.Total)

	if err != nil {
		return err
	}

	for _, item := range order.Items {
		q = "INSERT INTO order_item VALUES(?, ?, ?)"
		_, err := o.DB.Exec(q, order.Id, item.MenuId, item.Qty)

		if err != nil {
			return err
		}
	}

	return nil
}

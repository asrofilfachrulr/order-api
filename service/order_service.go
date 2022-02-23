package service

import (
	"database/sql"
	"errors"
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

func (o *OrderService) GetOrderById(id string) (*model.Order, []model.OrderItem, error.Error) {
	q := "SELECT * FROM order_list WHERE id = $1"
	r, err := o.DB.Query(q, id)
	if err != nil {
		return nil, nil, &error.InternalServerError{Err: err}
	}
	if !r.Next() {
		return nil, nil, &error.NotFoundError{Err: errors.New("order " + id + " not found!")}
	}

	order := model.Order{}
	err = r.Scan(&order.Id, &order.CreatedAt, &order.Status, &order.Total)
	if err != nil {
		return nil, nil, &error.InternalServerError{Err: err}
	}

	q = "SELECT * FROM order_item WHERE order_id = $1"
	r, err = o.DB.Query(q, order.Id)

	if err != nil {
		return nil, nil, &error.InternalServerError{Err: err}
	}
	defer r.Close()

	var orderItems []model.OrderItem

	for r.Next() {
		oi := model.OrderItem{}
		var temp string
		err = r.Scan(&temp, &oi.MenuId, &oi.Qty)
		if err != nil {
			return nil, nil, &error.InternalServerError{Err: err}
		}
		orderItems = append(orderItems, oi)
	}

	if len(orderItems) == 0 {
		return nil, nil, &error.InternalServerError{Err: errors.New("order item with order_id " + order.Id + " not found!")}
	}

	return &order, orderItems, nil
}

func (o *OrderService) UpdateOrderStatusById(id string) error.Error {
	q := "UPDATE order_list SET status = 'paid' WHERE id = $1"

	r, err := o.DB.Exec(q, id)
	if err != nil {
		return &error.InternalServerError{Err: err}
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return &error.NotFoundError{Err: errors.New("orderId " + id + " not found!")}
	}
	return nil
}

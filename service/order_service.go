package service

import (
	"database/sql"
	"orderapi/exception"
	"orderapi/model"

	"github.com/gin-gonic/gin"
)

type OrderService struct {
	DB *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		DB: db,
	}
}

func (o *OrderService) AddOrder(c *gin.Context, order *model.Order) {
	q := "INSERT INTO order_list VALUES(?, ?, ?, ?)"
	_, err := o.DB.Exec(q, order.Id, order.CreatedAt, order.Status, order.Total)

	if err != nil {
		exception.InternalServerError(c, err)
		return
	}

	for _, item := range order.Items {
		q = "INSERT INTO order_item VALUES(?, ?, ?)"
		_, err := o.DB.Exec(q, order.Id, item.MenuId, item.Qty)

		if err != nil {
			exception.InternalServerError(c, err)
			return
		}
	}
}

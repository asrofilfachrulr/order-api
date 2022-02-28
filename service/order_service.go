package service

import (
	"database/sql"
	"errors"
	"log"
	berror "orderapi/error"
	"orderapi/inmemory"
	"orderapi/model"
	"orderapi/utils"
	"time"
)

type OrderService struct {
	DB *sql.DB
}

func NewOrderService(db *sql.DB) *OrderService {
	return &OrderService{
		DB: db,
	}
}

func NullStringPointer() *string {
	var temp string = ""
	return &temp
}

func (o *OrderService) AddOrder(order *model.Order) berror.Error {
	q := "INSERT INTO order_list VALUES($1, $2, $3, $4, $5)"
	_, err := o.DB.Exec(q, order.Id, order.CreatedAt, order.UpdatedAt, order.Status, order.Total)

	if err != nil {
		return &berror.InternalServerError{Err: err}
	}

	for _, item := range order.Items {
		q = "INSERT INTO order_item VALUES($1, $2, $3)"
		_, err := o.DB.Exec(q, order.Id, item.Id, item.Qty)

		if err != nil {
			return &berror.InternalServerError{Err: err}
		}
	}
	return nil
}

func (o *OrderService) GetOrderById(id string) (*model.Order, berror.Error) {
	q := "SELECT * FROM order_list WHERE id = $1"
	r, err := o.DB.Query(q, id)
	if err != nil {
		return nil, &berror.InternalServerError{Err: err}
	}
	if !r.Next() {
		return nil, &berror.NotFoundError{Err: errors.New("order " + id + " not found!")}
	}

	order := model.Order{}
	err = r.Scan(&order.Id, &order.CreatedAt, &order.UpdatedAt, &order.Status, &order.Total)
	if err != nil {
		return nil, &berror.InternalServerError{Err: err}
	}

	q = "SELECT * FROM order_item WHERE order_id = $1"
	r, err = o.DB.Query(q, order.Id)

	if err != nil {
		return nil, &berror.InternalServerError{Err: err}
	}
	defer r.Close()

	var orderItems []model.OrderItem

	for r.Next() {
		oi := model.OrderItem{}

		err = r.Scan(NullStringPointer(), &oi.Id, &oi.Qty)
		if err != nil {
			return nil, &berror.InternalServerError{Err: err}
		}
		orderItems = append(orderItems, oi)
	}

	if len(orderItems) == 0 {
		return nil, &berror.InternalServerError{Err: errors.New("order item with order_id " + order.Id + " not found!")}
	}
	order.Items = orderItems

	return &order, nil
}

func (o *OrderService) UpdateOrderStatusById(id string, status string) berror.Error {
	q := "UPDATE order_list SET status = $1 WHERE id = $2"

	r, err := o.DB.Exec(q, status, id)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return &berror.NotFoundError{Err: errors.New("orderId " + id + " not found!")}
	}
	return nil
}

func (o *OrderService) UpdateOrderItemById(id string, oi []model.OrderItemUpdate) berror.Error {
	for _, item := range oi {
		// check for menuId existence in runtime
		if _, f := inmemory.ListMenuInmemory[item.Id]; !f {
			return &berror.NotFoundError{Err: errors.New("menu id " + utils.IntToString(item.Id) + " not found")}
		}
		// if qty is not present, it means user want to delete particular item
		if item.Qty == 0 {
			q := "DELETE FROM order_item WHERE order_id = $1 AND menu_id = $2"
			_, err := o.DB.Exec(q, id, item.Id)
			if err != nil {
				return &berror.InternalServerError{Err: err}
			}
		} else {
			// if menu is not exist, it means adding new item
			if e := o.CheckItemInOrder(id, item.Id); e != nil {
				oi := &model.OrderItemUpdate{
					Id:  item.Id,
					Qty: item.Qty,
				}
				err := o.AddItemForExistingOrder(id, oi)
				if err != nil {
					return err
				}
			} else {
				q := "UPDATE order_item SET qty = $1 WHERE order_id = $2 and menu_id = $3"
				_, err := o.DB.Exec(q, item.Qty, id, item.Id)
				if err != nil {
					return &berror.InternalServerError{Err: err}
				}
			}
		}
	}
	return nil
}

func (o *OrderService) AddItemForExistingOrder(id string, item *model.OrderItemUpdate) berror.Error {
	q := "INSERT INTO order_item VALUES($1, $2, $3)"
	_, err := o.DB.Exec(q, id, item.Id, item.Qty)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}
	return nil
}

func (o *OrderService) CheckOrderId(id string) berror.Error {
	q := "SELECT * FROM order_list WHERE id = $1"

	r, err := o.DB.Query(q, id)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}
	defer r.Close()

	if !r.Next() {
		return &berror.NotFoundError{Err: errors.New("order " + id + " not found")}
	}
	return nil
}
func (o *OrderService) CheckItemInOrder(orderId string, menuId int) berror.Error {
	q := "SELECT * FROM order_item WHERE order_id = $1 AND menu_id = $2"

	r, err := o.DB.Query(q, orderId, menuId)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}
	defer r.Close()

	if !r.Next() {
		return &berror.NotFoundError{Err: errors.New("order " + orderId + " not found")}
	}
	return nil
}

func (o *OrderService) UpdateStamp(id string) berror.Error {
	q := "UPDATE order_list SET updated_at = $1 WHERE id = $2"

	_, err := o.DB.Exec(q, time.Now(), id)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}
	return nil
}

func (o *OrderService) UpdateTotal(id string) berror.Error {
	// first, retrieve the list of item
	log.Printf("updating total for order id %s\n", id)
	q := "SELECT * FROM order_item WHERE order_id = $1"

	r, err := o.DB.Query(q, id)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}
	defer r.Close()

	oi := []model.OrderItemUpdate{}
	for r.Next() {
		log.Println("retrieve one item..")
		i := model.OrderItemUpdate{}
		var temp string
		err := r.Scan(&temp, &i.Id, &i.Qty)
		if err != nil {
			return &berror.InternalServerError{Err: err}
		}
		oi = append(oi, i)
	}

	log.Printf("retrieved items: %v\n", oi)
	// calculate the total from the retrieved list of item then set to order_list table
	total := 0
	for _, item := range oi {
		if it, f := inmemory.ListMenuInmemory[item.Id]; f {
			total += it.Price * int(item.Qty)
		} else {
			return &berror.InternalServerError{Err: errors.New("menu id " + utils.IntToString(item.Id) + " not found")}
		}
	}

	q = "UPDATE order_list SET total = $1 WHERE id = $2"
	_, err = o.DB.Exec(q, total, id)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}
	return nil
}

func (o *OrderService) DeleteOrderById(id string) berror.Error {
	// First deleting the order_item since they pointing to order_list
	q := "DELETE FROM order_item WHERE order_id = $1"
	_, err := o.DB.Exec(q, id)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}

	// Now deleting order_list
	q = "DELETE FROM order_list WHERE id = $1"
	_, err = o.DB.Exec(q, id)
	if err != nil {
		return &berror.InternalServerError{Err: err}
	}

	return nil
}

// no filter
func (o *OrderService) GetAllOrder(f *model.Filter, os *[]model.Order) berror.Error {
	q := ""
	var row *sql.Rows
	var err error
	today := time.Now()

	if f.Status != "" && f.From == "" {
		log.Println("case status isn't empty but time is")
		q = "SELECT * FROM order_list WHERE status = $1"
		row, err = o.DB.Query(q, f.Status)
	} else if f.Status == "" && f.From == "" {
		log.Println("case both status and time are empty")
		q = "SELECT * FROM order_list"
		row, err = o.DB.Query(q)
	} else if f.Status == "" && f.From != "" {
		log.Println("case status empty but time isnt empty (first look for 'from' value)")
		if today.Format("2006-01-02") == f.From {
			// if only today
			q = "SELECT * FROM order_list WHERE updated_at >= now()::date"
			row, err = o.DB.Query(q)
		} else {
			// if From is past
			q = "SELECT * FROM order_list WHERE updated_at BETWEEN $1::date AND $2::date + interval '1 day'"
			if f.To == "" {
				f.To = time.Now().Format("2006-01-02")
			}
			row, err = o.DB.Query(q, f.From, f.To)
		}
	} else {
		log.Println("case neither empty")
		q = "SELECT * FROM order_list WHERE updated_at BETWEEN $1::date AND $2::date + interval '1 day' AND status = $3"
		row, err = o.DB.Query(q, f.From, f.To, f.Status)
	}

	if err != nil {
		return &berror.InternalServerError{Err: err}
	}

	defer row.Close()

	for row.Next() {
		orderTemp := model.Order{}
		err := row.Scan(&orderTemp.Id, &orderTemp.CreatedAt, &orderTemp.UpdatedAt, &orderTemp.Status, &orderTemp.Total)
		if err != nil {
			return &berror.InternalServerError{Err: err}
		}
		*os = append(*os, orderTemp)
	}

	if len(*os) == 0 {
		return &berror.InternalServerError{Err: errors.New("DB is empty or query failed")}
	}

	return nil
}

func (o *OrderService) UpdateOrdersStatusByDate(f *model.Filter) berror.Error {
	var q string
	var err error

	if f.To == "" && f.From == "" {
		q = "UPDATE order_list SET status = $1"
		_, err = o.DB.Exec(q, f.Status)
	} else if f.From != "" {
		if f.To != "" {
			q = "UPDATE order_list SET status = $1 WHERE updated_at BETWEEN $2::date AND $3::date + interval '1 day'"
			_, err = o.DB.Exec(q, f.Status, f.From, f.To)
		} else {
			q = "UPDATE order_list SET status = $1 WHERE updated_at >= $2::date"
			_, err = o.DB.Exec(q, f.Status, f.From)
		}
	}

	if err != nil {
		return &berror.InternalServerError{Err: err}
	}

	return nil
}
func (o *OrderService) DeleteOrdersByFilter(f *model.Filter) berror.Error {
	var q string
	var err error

	if f.To == "" && f.From == "" {
		q = "DELETE FROM order_list SET status = $1"
		_, err = o.DB.Exec(q, f.Status)
	} else if f.From != "" {
		if f.To != "" {
			q = "DELETE FROM order_list SET status = $1 WHERE updated_at BETWEEN $2::date AND $3::date + interval '1 day'"
			_, err = o.DB.Exec(q, f.Status, f.From, f.To)
		} else {
			q = "DELETE FROM order_list SET status = $1 WHERE updated_at >= $2::date"
			_, err = o.DB.Exec(q, f.Status, f.From)
		}
	}

	if err != nil {
		return &berror.InternalServerError{Err: err}
	}

	return nil
}

package model

type OrderItem struct {
	Name string `json:"name" validate:"required"`
	Qty  int8   `json:"qty" validate:"required"`
}

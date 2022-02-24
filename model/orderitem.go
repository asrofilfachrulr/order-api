package model

type OrderItem struct {
	Id  int  `json:"id" validate:"required"`
	Qty int8 `json:"qty" validate:"required"`
}

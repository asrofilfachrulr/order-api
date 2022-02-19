package model

type OrderItem struct {
	MenuId int  `json:"menuId" validate:"required"`
	Qty    int8 `json:"qty" validate:"required"`
}

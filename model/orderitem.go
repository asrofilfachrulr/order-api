package model

type OrderItem struct {
	MenuId int  `json:"menuId"`
	Qty    int8 `json:"qty"`
}

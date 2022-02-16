package model

type OrderItem struct {
	MenuId string `json:"menuId"`
	Qty    int8   `json:"qty"`
}

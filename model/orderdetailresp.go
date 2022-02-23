package model

import "time"

type OrderDetailResp struct {
	Id        string                `json:"id"`
	Items     []OrderItemDetailResp `json:"items"`
	CreatedAt time.Time             `json:"createdAt"`
	UpdatedAt time.Time             `json:"updatedAt"`
	Status    string                `json:"status"`
	Total     int64                 `json:"total"`
}

type OrderItemDetailResp struct {
	Name string `json:"name"`
	Qty  int8   `json:"qty"`
}

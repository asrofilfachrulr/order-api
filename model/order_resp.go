package model

import "time"

type OrderResp struct {
	OrderId   string    `json:"orderId"`
	CreatedAt time.Time `json:"createdAt"`
	Total     int64     `json:"total"`
}

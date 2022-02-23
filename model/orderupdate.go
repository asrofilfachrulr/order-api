package model

type OrderUpdate struct {
	Status string      `json:"status"`
	Items  []OrderItem `json:"items"`
}

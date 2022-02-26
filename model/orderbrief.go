package model

type OrderBrief struct {
	Id        string `json:"id"`
	Status    string `json:"status"`
	UpdatedAt string `json:"updated_at"`
	Total     int64  `json:"total"`
}

package model

type CompleteOrder struct {
	Items  []Item
	Id     string
	Total  int
	IsPaid bool
	Code   string
}

type PaidStatus string

const (
	All    PaidStatus = "All"
	Paid   PaidStatus = "Paid"
	Unpaid PaidStatus = "Unpaid"
)

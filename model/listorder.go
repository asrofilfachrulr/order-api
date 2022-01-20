package model

type ListOrder struct {
	ListOrder []CompleteOrder
	IsPaid    bool
	Code      string
}

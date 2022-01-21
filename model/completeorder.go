package model

type CompleteOrder struct {
	Items  []Item
	Id     string
	Total  int
	IsPaid bool
	Code   string
}

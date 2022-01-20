package model

type CompleteOrder struct {
	Items []Item
	Id    string
	Total int
}

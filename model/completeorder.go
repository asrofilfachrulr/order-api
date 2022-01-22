package model

import "san_dong/dummy"

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

func (co *CompleteOrder) Pay(seat Seat) {
	// implement e-wallet api invokasition
	if seat.Status {
		if !dummy.Seat[seat.Pos] {
			dummy.Seat[seat.Pos] = true
		} else {

		}
	}
	co.IsPaid = true
}

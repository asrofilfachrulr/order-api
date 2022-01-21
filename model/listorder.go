package model

import "errors"

type ListOrder struct {
	// sync.Mutex
	ListOrder []CompleteOrder
}

func (lo *ListOrder) AppendOrder(order *CompleteOrder) error {
	if contain(*order, *lo) {
		return errors.New("Order id: " + order.Id + " existed")
	}
	lo.ListOrder = append(lo.ListOrder, *order)
	return nil
}

// tools for FinishOrder
func contain(order CompleteOrder, listorder ListOrder) bool {
	for _, lo := range listorder.ListOrder {
		if lo.Id == order.Id {
			return true
		}
	}
	return false
}

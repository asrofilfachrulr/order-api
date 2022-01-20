package controller

import (
	"errors"
	"math/rand"
	"san_dong/dummy"
	"san_dong/inmemory"
	"san_dong/model"
	"san_dong/validator"
	"strconv"
)

/*
*	Create a complete order object (CompleteOrder) that'll be used for payment
*	Front End sends []Item with each Item object present following attributes:
*   1. Qty <Present>
*   2. Food <Present>
*		a. Food.Id <Present>
*		b. Food.Name <Absent>
*		c. Food.Price <Absent>
*	3. Total <Absent>
*
*	MakeOrder completes all the absent attribute's values
*
*	If qty invalid => error
*
 */
func MakeOrder(item []model.Item) (*model.CompleteOrder, error) {
	t := 0
	err := validator.ValidateItem(item)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(item); i++ {

		item[i].Food.Name, _ = validator.ValidateMenuId(item[i].Food.Id)
		item[i].Food.Price = dummy.Price[item[i].Food.Id]
		item[i].Total = item[i].Food.Price * item[i].Qty
		t += item[i].Total
		// fmt.Println(t)
	}

	return &model.CompleteOrder{
		Items: item,
		Id:    "order-" + strconv.Itoa(rand.Int()),
		Total: t,
	}, nil
}

/*
*	Take created CompleteOrder send (add) to runtime object
*	which is ListOrder to provide shareable information
*	with cashier
*
 */

func contain(order model.CompleteOrder, listorder model.ListOrder) bool {
	for _, lo := range listorder.ListOrder {
		if lo.Id == order.Id {
			return true
		}
	}
	return false
}

func FinishOrder(order *model.CompleteOrder) error {
	if contain(*order, inmemory.ListOrderRuntime) {
		return errors.New("Order id: " + order.Id + " existed")
	}
	inmemory.ListOrderRuntime.ListOrder = append(inmemory.ListOrderRuntime.ListOrder, *order)

	return nil
}

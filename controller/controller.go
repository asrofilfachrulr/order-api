package controller

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

var menu = map[string]string{
	"mg001": "tempe goreng",
	"mk001": "bakso nanas",
	"mg002": "babi guling mentai",
	"mk002": "opor anjing manado",
}
var price = map[string]int{
	"mg001": 1000,
	"mk001": 12000,
	"mg002": 69000,
	"mk002": 99000,
}

type CompleteOrder struct {
	Items []Item
	Id    string
	Total int
}

type Item struct {
	Qty   int
	Food  Food
	Total int
}

type Food struct {
	Id    string
	Name  string
	Price int
}

/*
*	Front End sends []Item with each Item object present following attributes:
*   1. Qty <Present>
*   2. Food <Present>
*		a. Food.Id <Present>
*		b. Food.Name <Absent>
*		c. Food.Price <Absent>
*	3. Total <Absent>
*
*	MakeOrder completes all the absent attribute's values
 */
func MakeOrder(item *[]Item) *CompleteOrder {
	t := 0

	for i := 0; i < len(*item); i++ {
		(*item)[i].Food.Name, _ = SelectMenu((*item)[i].Food.Id)
		(*item)[i].Food.Price = price[(*item)[i].Food.Id]
		(*item)[i].Total = (*item)[i].Food.Price * (*item)[i].Qty
		t += (*item)[i].Total
		// fmt.Println(t)
	}

	fmt.Println(item)

	return &CompleteOrder{
		Items: *item,
		Id:    strconv.Itoa(rand.Int()),
		Total: t,
	}
}

func SelectMenu(id string) (string, error) {
	if v, ok := menu[id]; ok {
		return v, nil
	}

	return "", errors.New("menu doesnt exist")
}

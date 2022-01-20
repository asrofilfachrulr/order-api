package controller

import (
	"errors"
	"fmt"
	"math/rand"
	"san_dong/dummy"
	"san_dong/model"
	"san_dong/validator"
	"strconv"
)

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

		item[i].Food.Name, _ = SelectMenu(item[i].Food.Id)
		item[i].Food.Price = dummy.Price[item[i].Food.Id]
		item[i].Total = item[i].Food.Price * item[i].Qty
		t += item[i].Total
		// fmt.Println(t)
	}

	fmt.Println(item)

	return &model.CompleteOrder{
		Items: item,
		Id:    strconv.Itoa(rand.Int()),
		Total: t,
	}, nil
}

func SelectMenu(id string) (string, error) {
	if v, ok := dummy.Menu[id]; ok {
		return v, nil
	}

	return "", errors.New("menu doesnt exist")
}

package controller

import (
	"log"
	"san_dong/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectMakeOrder(t *testing.T) {
	var items []model.Item = []model.Item{
		{Qty: 2, Food: model.Food{Id: "mg001"}},
		{Qty: 3, Food: model.Food{Id: "mk001"}},
	}
	o, _ := MakeOrder(items)

	// Name of items should be correct
	assert.Equal(t, o.Items[0].Food.Name, "tempe goreng")
	assert.Equal(t, o.Items[1].Food.Name, "bakso nanas")

	// Item total should be correct
	assert.Equal(t, o.Items[0].Total, 2000)
	assert.Equal(t, o.Items[1].Total, 36000)

	// Order total should be correct
	assert.Equal(t, o.Total, 38000)

	// Have a proper Id
	assert.NotNil(t, o.Id)
}

func TestIncorrectOrder(t *testing.T) {
	var items []model.Item = []model.Item{
		{Qty: -1, Food: model.Food{Id: "mg001"}},              // invalid quantity
		{Qty: 3, Food: model.Food{Id: "iAm aN InvAlIId IDD"}}, // invalid id
	}
	_, err := MakeOrder(items)

	assert.EqualError(t, err, "error: invalid quantity (-1), invalid id (iAm aN InvAlIId IDD)")
}

func TestFinishOrder(t *testing.T) {
	var items []model.Item = []model.Item{
		{Qty: 2, Food: model.Food{Id: "mg001"}},
		{Qty: 3, Food: model.Food{Id: "mk001"}},
	}
	o, _ := MakeOrder(items)

	FinishOrder(o)
	log.Println(ListOrderRuntime)

	// ListOrderRuntime should have an item right now
	assert.Equal(t, len(ListOrderRuntime.ListOrder), 1)

	// the first item of ListOrderRuntime.ListOrder should be o
	assert.Equal(t, ListOrderRuntime.ListOrder[0], *o)
}

func TestMultipleFinishOrder(t *testing.T) {
	var items1 []model.Item = []model.Item{
		{Qty: 2, Food: model.Food{Id: "mg001"}},
		{Qty: 3, Food: model.Food{Id: "mk001"}},
	}
	order1, _ := MakeOrder(items1)
	FinishOrder(order1)

	var items2 []model.Item = []model.Item{
		{Qty: 1, Food: model.Food{Id: "mg002"}},
		{Qty: 1, Food: model.Food{Id: "mk002"}},
	}
	order2, _ := MakeOrder(items2)
	FinishOrder(order2)

	var items3 []model.Item = []model.Item{
		{Qty: 7, Food: model.Food{Id: "mg001"}},
	}
	order3, _ := MakeOrder(items3)
	FinishOrder(order3)

	log.Println(ListOrderRuntime)

	// ListOrderRuntime.ListOrder should have length of 3
	assert.Equal(t, len(ListOrderRuntime.ListOrder), 3)

	// ListOrderRuntime.ListOrder should have apropriate complete order
	assert.Equal(t, ListOrderRuntime.ListOrder[0], *order1)
	assert.Equal(t, ListOrderRuntime.ListOrder[1], *order2)
	assert.Equal(t, ListOrderRuntime.ListOrder[2], *order3)
}

func TestDuplicateIdFinishOrder(t *testing.T) {
	var items []model.Item = []model.Item{
		{Qty: 1, Food: model.Food{Id: "mg002"}},
		{Qty: 1, Food: model.Food{Id: "mk002"}},
	}

	co, _ := MakeOrder(items)
	first := FinishOrder(co)
	second := FinishOrder(co)

	assert.Nil(t, first)
	assert.EqualError(t, second, "Order id: "+co.Id+" existed")
}

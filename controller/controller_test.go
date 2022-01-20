package controller

import (
	"san_dong/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectExistedMenu(t *testing.T) {
	food, _ := SelectMenu("mg001")
	assert.Equal(t, food, "tempe goreng")
}

func TestSelectNotExistedMenu(t *testing.T) {
	_, err := SelectMenu("YouR DoG wiLl cOde")
	assert.Error(t, err, "menu doesnt exist")
}

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

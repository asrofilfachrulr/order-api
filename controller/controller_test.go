package controller

import (
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
	var items []Item = []Item{
		{Qty: 2, Food: Food{Id: "mg001"}},
		{Qty: 3, Food: Food{Id: "mk001"}},
	}
	o := MakeOrder(&items)

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

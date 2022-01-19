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

func TestSelectMultipleExistedMenu(t *testing.T) {
	foods := MultiSelectMenu("mg002", "mk002")
	assert.Equal(t, []string{"babi guling mentai", "opor anjing manado"}, foods)
}

func TestSelectMultipleMenuWithAnInvalidId(t *testing.T) {
	foods := MultiSelectMenu("mg001", "TONGKOL", "mk001")
	assert.Equal(t, foods[1], "")
}

package validator

import (
	"errors"
	"san_dong/dummy"
)

/* Valdiate whether menu's id is present and return its value */
func ValidateMenuId(id string) (string, error) {
	if v, ok := dummy.Menu[id]; ok {
		return v, nil
	}

	return "", errors.New("menu doesnt exist")
}

package validator

import (
	"errors"
	"san_dong/dummy"
	"san_dong/model"
	"strconv"
	"strings"
)

func ValidateItem(item []model.Item) error {
	var ErrMsg []string

	for _, it := range item {
		if it.Qty < 1 {
			msg := "invalid quantity (" + strconv.Itoa(it.Qty) + ")"
			ErrMsg = append(ErrMsg, msg)
		}
		if _, ok := dummy.Menu[it.Food.Id]; !ok {
			msg := "invalid id (" + it.Food.Id + ")"
			ErrMsg = append(ErrMsg, msg)
		}
	}
	if len(ErrMsg) > 0 {
		msg := "error: " + strings.Join(ErrMsg, ", ")
		return errors.New(msg)
	}
	return nil
}

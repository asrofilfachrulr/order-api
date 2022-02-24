package model

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"

	"github.com/go-playground/validator/v10"
)

type OrderUpdate struct {
	Status string      `json:"status" validate:"status"`
	Items  []OrderItem `json:"items" validate:"dive"`
}

func (o *OrderUpdate) ParseJSON(i io.ReadCloser) (OrderUpdate, error) {
	dataByte, err := ioutil.ReadAll(i)
	if err != nil {
		return OrderUpdate{}, err
	}

	newOrderUpdate := OrderUpdate{}

	err = json.Unmarshal(dataByte, &newOrderUpdate)
	if err != nil {
		return OrderUpdate{}, err
	}
	return newOrderUpdate, nil
}

func (o *OrderUpdate) ValidateMissing() error {
	missing := false

	if v := o.Status; v == "" {
		missing = true
	}
	if v := o.Items; len(v) == 0 {
		missing = true
	}

	if missing {
		return errors.New("attribute missing")
	}

	return nil
}

func ValidateStatus(f validator.FieldLevel) bool {
	status := f.Field().String()
	for _, s := range [...]string{"paid", "unpaid", "conflict"} {
		if s == status {
			return true
		}
	}
	return false
}

package model

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"

	"github.com/go-playground/validator/v10"
)

type OrderItemUpdate struct {
	Id  int  `json:"id" validate:"required"`
	Qty int8 `json:"qty"`
}

type OrderUpdate struct {
	Status string            `json:"status" validate:"status-validate"`
	Items  []OrderItemUpdate `json:"items" validate:"dive"`
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
	log.Printf("validating status, value %s\n", status)
	for _, s := range [...]string{"paid", "unpaid", "conflict"} {
		if s == status {
			log.Println("Found match")
			return true
		}
	}
	log.Println("Not found match")
	return false
}

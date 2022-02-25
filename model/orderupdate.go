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
	log.Printf("Success parse json: %v", newOrderUpdate)
	return newOrderUpdate, nil
}

func (o *OrderUpdate) CheckParsedJSON() error {
	missing := 0
	log.Println(o.Items)
	log.Println(len(o.Items))

	log.Printf("missing: %v\n", missing)
	if len(o.Items) == 0 {
		missing += 1
	}
	if o.Status == "" {
		missing += 1
	}

	if missing == 2 {
		return errors.New("please specifiy attribute that need to be changed")
	}

	return nil
}

func ValidateStatus(f validator.FieldLevel) bool {
	status := f.Field().String()
	for _, s := range [...]string{"paid", "unpaid", "conflict", ""} {
		if s == status {
			return true
		}
	}
	return false
}

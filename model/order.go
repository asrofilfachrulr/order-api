package model

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

type Order struct {
	Id        string      `json:"id"`
	Items     []OrderItem `json:"items" validate:"required,dive,required"`
	CreatedAt time.Time   `json:"createdAt"`
	Status    string      `json:"status"`
	Total     int64       `json:"total"`
}

func (o *Order) ParseJSON(i io.ReadCloser) (Order, error) {
	dataByte, err := ioutil.ReadAll(i)
	if err != nil {
		return Order{}, err
	}

	newOrder := Order{}

	err = json.Unmarshal(dataByte, &newOrder)
	if err != nil {
		return Order{}, err
	}
	return newOrder, nil
}

func ParseIdJSON(i io.ReadCloser) (map[string]string, error) {
	dataByte, err := ioutil.ReadAll(i)
	if err != nil {
		return nil, err
	}

	var body map[string]string

	err = json.Unmarshal(dataByte, &body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

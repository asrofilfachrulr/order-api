package model

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"time"
)

type Order struct {
	Id        string      `json:"id"`
	Items     []OrderItem `json:"items"`
	CreatedAt time.Time   `json:"createdAt"`
	Status    string      `json:"status"`
	Total     int64       `json:"total"`
}

func (o *Order) ParseJSON(i io.ReadCloser) (Order, error) {
	jsonByte, err := ioutil.ReadAll(i)
	if err != nil {
		return Order{}, err
	}

	newOrder := Order{}

	err = json.Unmarshal(jsonByte, &newOrder)
	if err != nil {
		return Order{}, err
	}

	return newOrder, nil
}

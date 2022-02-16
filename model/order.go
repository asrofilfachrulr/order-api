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

func (o *Order) ParseJSON(i io.ReadCloser) (*Order, error) {
	jsonByte, err := ioutil.ReadAll(i)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonByte, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

package model

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"time"
)

type Filter struct {
	Status string
	From   string
	To     string
}

func (f *Filter) ParseJSON(i io.ReadCloser) (Filter, error) {
	dataByte, err := ioutil.ReadAll(i)
	if err != nil {
		return Filter{}, err
	}

	newFilter := Filter{}

	err = json.Unmarshal(dataByte, &newFilter)
	if err != nil {
		return Filter{}, err
	}
	return newFilter, nil
}

func (f *Filter) ValidateStatus() {
	if f.Status != "unpaid" && f.Status != "paid" && f.Status != "conflict" {
		f.Status = ""
	}
}

func (f *Filter) ValidateTime() {
	if f.From == "" {
		return
	}
	from, err := time.Parse("2006-01-02", f.From)
	if err != nil {
		log.Println("error parsing from")
		log.Println(err)
		f.From = ""
		return
	}

	if f.To == "" {
		return
	}
	to, err := time.Parse("2006-01-02", f.To)
	if err != nil {
		log.Println("error parsing to")
		f.To = ""
		return
	}

	if from.After(to) {
		log.Println("illogical address")

		f.From = ""
		f.To = ""
	}

}

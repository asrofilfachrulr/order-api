package model

import (
	"time"
)

type Filter struct {
	Status string
	From   string
	To     string
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
	from, err := time.Parse("2016-02-01", f.From)
	if err != nil {
		f.From = ""
		return
	}

	if f.To == "" {
		return
	}
	to, err := time.Parse("2016-02-01", f.To)
	if err != nil {
		f.To = ""
		return
	}

	if from.After(to) {
		f.From = ""
		f.To = ""
	}

}

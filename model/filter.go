package model

import (
	"log"
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

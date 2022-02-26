package model

import "time"

type Filter struct {
	Status string
	Time   time.Time
}

func (f *Filter) ValidateStatus() {
	if f.Status != "unpaid" && f.Status != "paid" && f.Status != "conflict" {
		f.Status = ""
	}
}

func (f *Filter) ValidateTime() {
	if f.Time.After(time.Now()) {
		f.Time = time.Time{}
	}
}

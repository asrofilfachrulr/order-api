package model

type Menu struct {
	Id    int    `sql:"id"`
	Price int    `sql:"price"`
	Name  string `sql:"name"`
}

package service

import (
	"database/sql"
	"orderapi/inmemory"
	"orderapi/model"
)

type MenuService struct {
	DB *sql.DB
}

func NewMenuService(db *sql.DB) *MenuService {
	return &MenuService{
		DB: db,
	}
}

func (ms *MenuService) LoadMenu() error {
	q := "SELECT * FROM menu_list"

	rows, err := ms.DB.Query(q)

	if err != nil {
		return err
	}

	for rows.Next() {
		m := model.Menu{}
		var temp string

		err := rows.Scan(&m.Id, &temp, &m.Price)

		if err != nil {
			return err
		}

		inmemory.ListMenuInmemory[m.Id] = m.Price
	}
	return nil
}

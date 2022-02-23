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

		err := rows.Scan(&m.Id, &m.Name, &m.Price)

		if err != nil {
			return err
		}

		inmemory.ListMenuInmemory[m.Name] = &inmemory.MenuInfoByName{
			Price: m.Price,
			Id:    m.Id,
		}
		delete(inmemory.ListMenuInmemory, "")

		inmemory.MenuIdNameMap[m.Id] = m.Name
	}
	return nil
}

package service

import (
	"database/sql"
	"orderapi/inmemory"
	"orderapi/model"
)

type Service struct {
	DB *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) LoadMenu() error {
	q := "SELECT * FROM menu_list"

	rows, err := s.DB.Query(q)

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

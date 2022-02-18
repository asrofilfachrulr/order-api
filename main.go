package main

import (
	"fmt"
	"orderapi/app"
	"orderapi/controller"
	"orderapi/handler"
	"orderapi/inmemory"
	"orderapi/service"
	"orderapi/util"
)

func main() {
	db := app.NewDB()
	s := service.NewService(db)
	c := controller.NewController(s)
	h := handler.NewHandler(c)
	r := app.NewRouter(h)

	err := s.LoadMenu()
	util.ExitIfError(err)

	fmt.Println(inmemory.ListMenuInmemory)

	r.Run()
}

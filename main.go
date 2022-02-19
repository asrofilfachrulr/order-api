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

	m := service.NewMenuService(db)
	err := m.LoadMenu()
	util.ExitIfError(err)

	s := service.NewOrderService(db)
	c := controller.NewController(s)
	h := handler.NewHandler(c)
	r := app.NewRouter(h)

	fmt.Println(inmemory.ListMenuInmemory)

	r.Run()
}

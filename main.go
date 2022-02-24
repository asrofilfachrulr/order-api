package main

import (
	"fmt"
	"orderapi/app"
	"orderapi/controller"
	"orderapi/exception"
	"orderapi/handler"
	"orderapi/inmemory"
	"orderapi/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	// db connectionP
	db := app.NewDB()

	// service
	m := service.NewMenuService(db.DB)
	err := m.LoadMenu()
	exception.TerminateRuntimeIfError(err)

	s := service.NewOrderService(db.DB)

	// controller
	c := controller.NewController(s)
	// validator
	v := validator.New()
	// handler
	h := handler.NewHandler(c, v)
	// router
	r := app.NewRouter(h)

	fmt.Println("Load Menu..")
	for i, item := range inmemory.ListMenuInmemory {
		fmt.Printf("%d\t%s\t%d\n", i, item.Name, item.Price)
	}
	r.Run()
}

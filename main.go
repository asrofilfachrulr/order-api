package main

import (
	"orderapi/app"
	"orderapi/controller"
	"orderapi/exception"
	"orderapi/handler"
	"orderapi/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()

	m := service.NewMenuService(db)
	err := m.LoadMenu()
	exception.TerminateRuntimeIfError(err)

	s := service.NewOrderService(db)
	c := controller.NewController(s)
	v := validator.New()
	h := handler.NewHandler(c, v)
	r := app.NewRouter(h)

	// fmt.Println(inmemory.ListMenuInmemory)

	r.Run()
}

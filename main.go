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
	_ "github.com/go-sql-driver/mysql"
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

	fmt.Println(inmemory.ListMenuInmemory)

	r.Run()
}

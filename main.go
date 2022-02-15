package main

import (
	"orderapi/app"
	"orderapi/handler"
)

func main() {
	h := handler.NewHandler()
	r := app.NewRouter(h)

	r.Run()
}

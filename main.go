package main

import (
	"fmt"
	"net/http"
	"quoteapp/controller"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/routes"
)

func main() {
	db := db.NewDB(
		"localhost",
		"postgres",
		"Kmzway87a@",
		"quoteapp",
		5432,
	)

	db.Migrate()
	conn := db.DB()

	model := model.NewQuoteModel(conn)
	controller := controller.NewQuoteController(model)

	//edit or alter these code to include user model and controller

	router := routes.NewRoute(controller)

	fmt.Println("starting api server at http://localhost:8080")
	http.ListenAndServe(":8080", router.Run())
}

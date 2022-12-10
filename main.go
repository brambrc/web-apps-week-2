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
		"asd123",
		"data",
		5432,
	)

	db.Migrate()
	conn := db.DB()

	modelQuote := model.NewQuoteModel(conn)
	modelUser := model.NewUsersModel(conn)
	quote := controller.NewQuoteController(modelQuote)
	user := controller.NewUsersController(modelUser)

	router := routes.NewRoute(quote, user)

	fmt.Println("starting api server at http://localhost:8080")
	http.ListenAndServe(":8080", router.Run())
}

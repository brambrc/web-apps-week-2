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
		"ga123124",     //password
		"anime_quotes", //nama db
		5432,
	)

	db.Migrate()
	conn := db.DB()

	modelQuotes := model.NewQuoteModel(conn)
	modelUsers := model.NewUsersModel(conn)

	controllerQuotes := controller.NewQuoteController(modelQuotes)

	//edit or alter these code to include user model and controller
	controllerUsers := controller.NewUsersController(modelUsers)

	router := routes.NewRoute(controllerQuotes, controllerUsers)

	fmt.Println("starting api server at http://localhost:8080")
	http.ListenAndServe(":8080", router.Run())
}

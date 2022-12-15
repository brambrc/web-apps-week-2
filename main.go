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
		"warva14",
		"test",
		5432,
	)

	db.Migrate()
	conn := db.DB()

	quoteModel := model.NewQuoteModel(conn)
	quoteMontroller := controller.NewQuoteController(quoteModel)

	usermodel := model.NewUsersModel(conn)
	usercontroller := controller.NewUsersController(usermodel)

	//edit or alter these code to include user model and controller

	router := routes.NewRoute(quoteMontroller, usercontroller)

	fmt.Println("starting api server at http://localhost:8080")
	http.ListenAndServe(":8080", router.Run())
}

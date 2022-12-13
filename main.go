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
		"192.168.0.102",
		"postgres",
		"password",
		"quoteapp",
		5432,
	)

	db.Migrate()
	conn := db.DB()

	qm := model.NewQuoteModel(conn)
	qoute := controller.NewQuoteController(qm)

	usermodel := model.NewUsersModel(conn)
	usercontroller := controller.NewUsersController(usermodel)

	router := routes.NewRoute(qoute, usercontroller)

	fmt.Println("starting api server at http://localhost:8080")
	panic(http.ListenAndServe(":8080", router.Run()))
}

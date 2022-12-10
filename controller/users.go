package controller

import (
	"encoding/json"
	"net/http"
	"quoteapp/model"
	"quoteapp/view"
)

type Users struct {
	users *model.Users
}

// func NewUsersController(q *model.Users) *Users {
// 	return &Users{users: q}
// }

func (q *Users) Store(w http.ResponseWriter, r *http.Request) {

	// add code to hit create model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) FindAll(w http.ResponseWriter, r *http.Request) {

	// add code to hit find all model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: users})
}

func (q *Users) FindByID(w http.ResponseWriter, r *http.Request) {

	// add code to hit find by id model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: user})
}

func (q *Users) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// add code to hit update model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// add code to hit delete model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

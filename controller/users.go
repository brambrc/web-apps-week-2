package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"quoteapp/db"
	"quoteapp/model"
	"quoteapp/view"
)

type Users struct {
	users *model.Users
}

func NewUsersController(q *model.Users) *Users {
	return &Users{users: q}
}

func (q *Users) Store(w http.ResponseWriter, r *http.Request) {

	// add code to hit create model user here

	body, err := io.ReadAll(r.Body)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	user := db.Users{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	err = q.users.Create(&user)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]bool{
		"add": true,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) FindAll(w http.ResponseWriter, r *http.Request) {

	// add code to hit find all model user here

	data, err := q.users.FindAll()
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) FindByID(w http.ResponseWriter, r *http.Request) {

	// add code to hit find by id model user here

	id := r.URL.Query().Get("id")

	data, err := q.users.FindByID(id)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) Update(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	user := db.Users{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}
	// add code to hit update model user here
	er := q.users.Update(id, user)
	if er != nil {
		view.ErrorRespond(w, er)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Error: er})
}

func (q *Users) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// add code to hit delete model user here
	err := q.users.Delete(id)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Error: err})
}

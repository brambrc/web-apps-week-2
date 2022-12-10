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

	data := map[string]string{
		"add": "berhasil menambah data user",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) FindAll(w http.ResponseWriter, r *http.Request) {

	// add code to hit find all model user here
	users, err := q.users.FindAll()
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: users})
}

func (q *Users) FindByID(w http.ResponseWriter, r *http.Request) {

	// add code to hit find by id model user here
	id := r.URL.Query().Get("id")

	user, err := q.users.FindByID(id)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: user})
}

func (q *Users) Update(w http.ResponseWriter, r *http.Request) {
	// add code to hit update model user here
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

	err = q.users.Update(id, user)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]string{
		"update": "berhasil mengupdate data user",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	// add code to hit delete model user here
	err := q.users.Delete(id)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]string{
		"delete": "berhasil menghapus data user",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

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
	x, err := io.ReadAll(r.Body)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	y := db.Users{}
	err = json.Unmarshal(x, &y)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	err = q.users.Create(&y)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]bool{
		"add": true,
	}
	// add code to hit create model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) FindAll(w http.ResponseWriter, r *http.Request) {
	data := []db.Users{}
	if _, err := q.users.FindAll(); err != nil {
		view.ErrorRespond(w, err)
		return
	}
	// add code to hit find all model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) FindByID(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query().Get("id")
	if _, err := q.users.FindByID(data); err != nil {
		view.ErrorRespond(w, err)
		return
	}
	// add code to hit find by id model user here

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

	err = q.users.Update(id, &user)
	if err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]string{
		"pesan": "berhasil mengupdate user dengan id : " + id,
	}

	// add code to hit update model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

func (q *Users) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if err := q.users.Delete(id); err != nil {
		view.ErrorRespond(w, err)
		return
	}

	data := map[string]string{
		"pesan": "berhasil menghapus user dengan id : " + id,
	}
	// add code to hit delete model user here

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(view.Response{Data: data})
}

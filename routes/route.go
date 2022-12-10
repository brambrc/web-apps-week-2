package routes

import (
	"net/http"
	"quoteapp/controller"
)

type Route struct {
	mux   *http.ServeMux
	quote *controller.Quote
	users *controller.Users
}

func NewRoute(q *controller.Quote, u *controller.Users) *Route {
	route := &Route{
		mux:   http.NewServeMux(),
		quote: q,
		users: u,
	}

	route.routes()
	return route
}

func (r *Route) Run() *http.ServeMux {
	return r.mux
}

func (r *Route) routes() {
	r.mux.Handle("/fetch", r.get(http.HandlerFunc(r.quote.Fetch)))
	r.mux.Handle("/select", r.get(http.HandlerFunc(r.quote.Show)))
	r.mux.Handle("/count", r.get(http.HandlerFunc(r.quote.Count)))
	r.mux.Handle("/add", r.post(http.HandlerFunc(r.quote.Store)))

	//add more routes for users here
	r.mux.Handle("/users/store", r.post(http.HandlerFunc(r.users.Store)))
	r.mux.Handle("/users/findall", r.get(http.HandlerFunc(r.users.FindAll)))
	r.mux.Handle("/users/findbyid", r.get(http.HandlerFunc(r.users.FindByID)))
	r.mux.Handle("/users/update", r.post(http.HandlerFunc(r.users.Update)))
	r.mux.Handle("/users/delete", r.post(http.HandlerFunc(r.users.Delete)))

}

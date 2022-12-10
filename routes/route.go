package routes

import (
	"net/http"
	"quoteapp/controller"
)

type Route struct {
	mux   *http.ServeMux
	quote *controller.Quote
	user  *controller.Users
}

func NewRoute(q *controller.Quote, u *controller.Users) *Route {
	route := &Route{
		mux:   http.NewServeMux(),
		quote: q,
		user:  u,
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

	r.mux.Handle("/store", r.post(http.HandlerFunc(r.user.Store)))
	r.mux.Handle("/findall", r.get(http.HandlerFunc(r.user.FindAll)))
	r.mux.Handle("/findbyid", r.get(http.HandlerFunc(r.user.FindByID)))
	r.mux.Handle("/update", r.post(http.HandlerFunc(r.user.Update)))
	r.mux.Handle("/delete", r.get(http.HandlerFunc(r.user.Delete)))
	//add more routes for users here
}

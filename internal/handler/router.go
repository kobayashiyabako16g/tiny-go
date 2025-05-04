package handler

import "net/http"

type Router struct {
	uh User
}

func NewRouter(uh User) Router {
	return Router{uh}
}

func (r *Router) HandleRequest(mux *http.ServeMux) {
	mux.Handle("/users/{id}", r.uh.GetUser())
}

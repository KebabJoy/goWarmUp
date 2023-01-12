package routes

import (
	v1 "Application/internal/api/v1/routes"
	"Application/internal/conns"
	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
	conns  *conns.Conns
}

func New(router *mux.Router, conns *conns.Conns) *Router {
	return &Router{router, conns}
}

func (r *Router) Setup() {
	v1Subrouter := v1.NewSubrouter(r.router, r.conns)
	v1Subrouter.Setup()
}

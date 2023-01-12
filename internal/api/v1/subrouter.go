package v1

import (
	"VSpace/internal/api/v1/users"
	"VSpace/internal/conns"
	"github.com/gorilla/mux"
	"net/http"
)

type Subrouter struct {
	router *mux.Router
	conns  *conns.Conns
}

func NewSubrouter(router *mux.Router, conns *conns.Conns) *Subrouter {
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	return &Subrouter{subrouter, conns}
}

func (s *Subrouter) SetupRoutes() {
	usersController := users.Loader{MainDb: s.conns.MainDB}
	s.router.HandleFunc("/users", usersController.GetUsers).Methods(http.MethodGet)
}

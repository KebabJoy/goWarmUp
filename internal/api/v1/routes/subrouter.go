package routes

import (
	"Application/internal/api/v1/files"
	"Application/internal/conns"
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

func (s *Subrouter) Setup() {
	zipController := files.Controller{MainDb: s.conns.MainDB}
	s.router.HandleFunc("/zip", zipController.Create).Methods(http.MethodPost)
	s.router.HandleFunc("/files", zipController.Index).Methods(http.MethodGet)
	s.router.HandleFunc("/files/{filename}", zipController.Show).Methods(http.MethodGet)
}

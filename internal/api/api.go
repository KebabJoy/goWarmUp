package api

import (
	"Application/internal/api/routes"
	"Application/internal/conns"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.uber.org/zap"
	"net/http"
)

type Server interface {
	Config()
}

type Api struct {
	conns  *conns.Conns
	logger *zap.Logger
}

func New(conns *conns.Conns, logger *zap.Logger) Server {
	return &Api{conns, logger}
}

func (a *Api) Config() {
	muxRouter := mux.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodPost,
			http.MethodGet,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})
	apiRoutes := routes.New(muxRouter, a.conns)
	apiRoutes.Setup()
	handler := cors.Handler(muxRouter)

	err := http.ListenAndServe(":3000", handler)

	if err != nil {
		a.logger.Log(zap.ErrorLevel, "FAIL ON RUNNING HTTP")
	}
}

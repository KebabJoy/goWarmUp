package api

import (
	"Application/internal/api/routes"
	"Application/internal/conns"
	"github.com/gorilla/mux"
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
	apiRoutes := routes.New(muxRouter, a.conns)
	apiRoutes.Setup()
	err := http.ListenAndServe(":3000", muxRouter)

	if err != nil {
		a.logger.Log(zap.ErrorLevel, "FAIL ON RUNNING HTTP")
	}
}

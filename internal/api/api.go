package api

import (
	"VSpace/internal/conns"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type Api struct {
	conns  *conns.Conns
	router *mux.Router
	logger *zap.Logger
}

func New(conns *conns.Conns, router *mux.Router, logger *zap.Logger) *Api {
	return &Api{conns, router, logger}
}

func (a *Api) ConfigServer() {
	a.SetupRoutes()
	err := http.ListenAndServe(":3000", a.router)

	if err != nil {
		a.logger.Log(zap.ErrorLevel, "FAIL ON RUNNING HTTP")
	}

}

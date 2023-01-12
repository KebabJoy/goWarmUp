package api

import v1 "VSpace/internal/api/v1"

type Routes interface {
	Setup()
}

func (a *Api) SetupRoutes() {
	v1Subrouter := v1.NewSubrouter(a.router, a.conns)
	v1Subrouter.SetupRoutes()
}

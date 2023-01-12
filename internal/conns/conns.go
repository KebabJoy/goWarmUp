package conns

import (
	"database/sql"
)

type Conns struct {
	MainDB *sql.DB
	//Cache  *cache.Cache
}

func New(mainDB *sql.DB) *Conns {
	return &Conns{mainDB}
}

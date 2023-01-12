package conns

import "github.com/jmoiron/sqlx"

type Conns struct {
	MainDB *sqlx.DB
	//Cache  *cache.Cache
}

func New(mainDB *sqlx.DB) *Conns {
	return &Conns{mainDB}
}

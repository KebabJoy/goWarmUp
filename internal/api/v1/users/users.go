package users

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

const dbQuery = `SELECT
name,
email
FROM users
WHERE email = ? AND password = ?`

type Users struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
	//roleId   int16
}

type Loader struct {
	MainDb *sqlx.DB
	//cache *cache.Cache
}

func (l *Loader) loadUsers() []Users {
	return []Users{{"lol", "lol", "lol"}, {"lol", "lel", "jop"}}
}

func (l *Loader) GetUsers(w http.ResponseWriter, req *http.Request) {
	fmt.Println(l.loadUsers())
	json.NewEncoder(w).Encode(l.loadUsers())
}

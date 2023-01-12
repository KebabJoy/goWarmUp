package files

import (
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type Controller struct {
	MainDb *sqlx.DB
}

func (c *Controller) Create(w http.ResponseWriter, req *http.Request) {
	fmt.Println(c.loadFiles())
	req.ParseMultipartForm(maxFileSize)

	file, _, err := req.FormFile("myFile")
	if err != nil {
		// На клиент кнш не круто отдавать такие сообщения, это я для дебага через постман сделал
		json.NewEncoder(w).Encode(failedResponse(err.Error()))
		return
	}

	c.fetchLastId()
	fmt.Println("SUCCESS: ", file)

	json.NewEncoder(w).Encode(c.loadFiles())
}

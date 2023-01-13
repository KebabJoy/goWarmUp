package files

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	maxFileSize = 1e10
	// success добавляли изначально, думал буду делать асинхронную архивацию файлов.
	// Чтобы потом по ID джобы можно было получить файл либо узнать, что еще не запроцессилось. Но не хочется много времени на тестовое тратить..
	dbQuery = `SELECT
				id,
				success,
				filename
				FROM zip_logs`
)

type Record struct {
	Id       int    `db:"id"`
	Success  int    `db:"success"`
	Filename string `db:"filename"`
}

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Filename string `json:"filename,omitempty"`
}

func (c *Controller) loadFiles() ([]Record, error) {
	// Сори, без пагинации(
	var files []Record
	err := c.MainDb.Select(&files, dbQuery)
	if err != nil {
		fmt.Println("ERROR ON FETCHING FILES QUERY", err)
		return []Record{}, err
	}

	return files, nil
}

func failedResponse(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(Response{Success: false, Message: msg})
}

func (c *Controller) recordExists(filename string) (bool, error) {
	rows, err := c.MainDb.Queryx(dbQuery)
	defer rows.Close()
	if err != nil {
		fmt.Println("ERROR ON FETCHING LAST ID QUERY", err)
		return false, err
	}

	var rec Record
	for rows.Next() {
		if err = rows.StructScan(&rec); err != nil {
			fmt.Println("ERROR ON SCANNING LAST ID RECORD", err)
			return false, err
		}
		fmt.Println(rec)
	}

	return false, nil
}

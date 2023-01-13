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
	dbQueryIndex = `SELECT
				id,
				success,
				filename
				FROM zip_logs`
	dbQueryExists = `SELECT COUNT(1) FROM zip_logs WHERE filename=?`
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
	err := c.MainDb.Select(&files, dbQueryIndex)
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
	var check bool
	err := c.MainDb.Get(&check, dbQueryExists, filename)
	if err != nil {
		fmt.Println("ERROR ON FETCHING EXISTS QUERY", err)
		return false, err
	}

	return true, nil
}

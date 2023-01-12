package files

import (
	"fmt"
)

const (
	maxFileSize = 1e10
	dbQuery     = `SELECT
				id,
				success
				FROM zip_logs`
)

type Record struct {
	Success  int `db:"success"`
	Filename int `db:"filename"`
}

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Filename string `json:"filename,omitempty"`
}

func (c *Controller) loadFiles() []Record {
	return []Record{}
}

func failedResponse(msg string) Response {
	return Response{Success: false, Message: msg}
}

func (c *Controller) fetchLastRecord(filename string) (*Record, error) {
	rows, err := c.MainDb.Queryx(dbQuery + " ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Println("ERROR ON FETCHING LAST ID QUERY", err)
		return &Record{}, err
	}

	var rec Record
	for rows.Next() {
		if err = rows.StructScan(&rec); err != nil {
			fmt.Println("ERROR ON SCANNING LAST ID RECORD", err)
			return &Record{}, err
		}
		fmt.Println(rec)
	}

	return &rec, nil
}

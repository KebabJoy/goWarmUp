package files

import "fmt"

const (
	maxFileSize = 1e10
	dbQuery     = `SELECT
				id,
				success
				FROM zip_logs`
)

type Record struct {
	ID      int `db:"id"`
	Success int `db:"success"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	JobId   string `json:"job_id,omitempty"`
}

func (c *Controller) loadFiles() []Record {
	return []Record{}
}

func failedResponse(msg string) Response {
	return Response{Success: false, Message: msg}
}

func (c *Controller) fetchLastId() (int, error) {
	rows, err := c.MainDb.Queryx(dbQuery + " ORDER BY id DESC LIMIT 1")
	if err != nil {
		fmt.Println("ERROR ON FETCHING LAST ID QUERY", err)
		return 0, err
	}

	for rows.Next() {
		var rec Record
		if err = rows.StructScan(&rec); err != nil {
			fmt.Println("ERROR ON SCANNING LAST ID RECORD", err)
			return 0, err
		}
		fmt.Println(rec)
	}

	return 0, nil
}

package files

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"net/http"
	"strings"
)

type Controller struct {
	MainDb *sqlx.DB
}

func (c *Controller) Create(w http.ResponseWriter, req *http.Request) {
	fmt.Println(c.loadFiles())
	req.ParseMultipartForm(maxFileSize)

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		// На клиент кнш не круто отдавать такие сообщения, это я для дебага через постман сделал
		json.NewEncoder(w).Encode(failedResponse(err.Error()))
		return
	}

	arrayFilename := strings.Split(handler.Filename, ".")
	format := arrayFilename[len(arrayFilename)-1]

	tempFile, err := ioutil.TempFile("./temp", "upload-*."+format)
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("ERROR saving tempfile", err)
	}
	tempFile.Write(fileBytes)

	archiveName := uuid.New().String()
	err = c.buildZip(archiveName, tempFile.Name())
	if err != nil {
		json.NewEncoder(w).Encode(failedResponse(err.Error()))
		return
	}
	fmt.Println("SUCCESS: ", file)

	response := Response{true, "", archiveName}
	json.NewEncoder(w).Encode(response)
}

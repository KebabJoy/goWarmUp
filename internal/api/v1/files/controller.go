package files

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Controller struct {
	MainDb *sqlx.DB
}

func (c *Controller) Show(w http.ResponseWriter, req *http.Request) {
	filename := mux.Vars(req)["filename"]
	if check, err := c.recordExists(filename); !check || err != nil {
		failedResponse(w, "Param `filename` is missing or has invalid value")
		return
	}

	// Пускай 500 будет, если файла нет. Выше проверка по логам есть. Файлы в AWS персистентно лежат
	file, err := os.Open(filename + fileFormat)
	if err != nil {
		panic(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		failedResponse(w, err.Error())
		return
	}
	fileHeader := make([]byte, 512)
	_, err = file.Read(fileHeader)
	if err != nil {
		failedResponse(w, err.Error())
		return
	}
	requestRange := req.Header.Get("range")
	if requestRange == "" {
		w.Header().Set("Content-Length", strconv.Itoa(int(fileInfo.Size())))
		file.Seek(0, 0)
		io.Copy(w, file)
		return
	}

	requestRange = requestRange[6:]
	splitRange := strings.Split(requestRange, "-")
	if len(splitRange) != 2 {
		failedResponse(w, "invalid values for header 'Range'")
		return
	}
	begin, err := strconv.ParseInt(splitRange[0], 10, 64)
	if err != nil {
		failedResponse(w, err.Error())
		return
	}
	end, err := strconv.ParseInt(splitRange[1], 10, 64)
	if err != nil {
		failedResponse(w, err.Error())
		return
	}
	if begin > fileInfo.Size() || end > fileInfo.Size() {
		failedResponse(w, "range out of bounds for file")
		return
	}
	if begin >= end {
		failedResponse(w, "range begin cannot be bigger than range end")
		return
	}
	w.Header().Set("Content-Length", strconv.FormatInt(end-begin+1, 10))
	w.Header().Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", begin, end, fileInfo.Size()))
	w.WriteHeader(http.StatusPartialContent)
	file.Seek(begin, 0)
	io.CopyN(w, file, end-begin)
}

func (c *Controller) Index(w http.ResponseWriter, _ *http.Request) {
	files, err := c.loadFiles()
	if err != nil {
		failedResponse(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(files)
}

func (c *Controller) Create(w http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(maxFileSize)

	file, handler, err := req.FormFile("myFile")
	if err != nil {
		// На клиент кнш не круто отдавать такие сообщения, это я для дебага через постман сделал
		failedResponse(w, err.Error())
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
		failedResponse(w, err.Error())
		return
	}
	fmt.Println("SUCCESS: ", file)

	response := Response{true, "", archiveName}
	json.NewEncoder(w).Encode(response)
}

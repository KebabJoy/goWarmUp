package files

import (
	"fmt"
	"os"
	"strconv"
)

func buildZip(fileId int, format string) {
	archive, err := os.Create("archive.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	//zipWriter := zip.NewWriter(archive)

	fmt.Println("opening first file...")
	f1, err := os.Open("uploaded-" + strconv.Itoa(fileId) + "." + format)
	if err != nil {
		panic(err)
	}
	defer f1.Close()
}

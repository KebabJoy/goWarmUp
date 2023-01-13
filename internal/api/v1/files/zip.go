package files

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

const (
	fileFormat = ".zip"
)

func (c *Controller) buildZip(archiveName string, filename string) error {
	archive, err := os.Create(archiveName + fileFormat)
	if err != nil {
		fmt.Println("ERROR creating archive", err)
		return err
	}
	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	f1, err := os.Open(filename)
	defer f1.Close()
	if err != nil {
		fmt.Println("ERROR opening first file...")
		return err
	}

	w1, err := zipWriter.Create(filename)
	if err != nil {
		fmt.Println("ERROR creating first file...")
		return err
	}
	if _, err := io.Copy(w1, f1); err != nil {
		fmt.Println("ERROR copying first file...")
		return err
	}

	c.MainDb.Queryx("INSERT INTO zip_logs (success, filename) VALUES (TRUE, ?)", archiveName)
	return nil
}

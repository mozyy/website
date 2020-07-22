package util

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const distDir = "dist/"

// Compress is compress directive to dist path (sources => name.zip/dist/sources)
func Compress(basePath, source, name string) (string, error) {
	zipPath := filepath.Join(basePath, name+".zip")
	file, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a new zip archive.
	w := zip.NewWriter(file)

	// Add some files to the archive.
	err = addFiles(w, filepath.Join(basePath, source), distDir)

	if err != nil {
		return "", err
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		return "", err
	}
	fmt.Printf("压缩成功: %s\n", zipPath)
	return zipPath, err
}

func addFiles(w *zip.Writer, basePath, baseInZip string) error {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return err
	}

	for _, file := range files {
		newPath := filepath.Join(basePath, file.Name())
		newInZip := filepath.Join(baseInZip + file.Name())
		// fmt.Printf("add file: newPath: %s, newInZip: %s, basePath: %s, baseZip: %s\n", newPath, newInZip, basePath, baseInZip)
		fileHeader:=&zip.FileHeader{
			Name: newInZip,
			Modified: time.Now(),
			Method: zip.Deflate,
		}
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(newPath)
			if err != nil {
				return err
			}
			// Add some files to the archive.
			// f, err := w.Create(newInZip)
			// f, err := w.CreateHeader(&zip.FileHeader{
			// 	Name: newInZip,
			// 	Modified: time.Now(),
			// })
			// if err != nil {
			// 	return err
			// }
			f, err := w.CreateHeader(fileHeader)
			if err != nil {
				return err
			}
			_, err = f.Write(dat)
			if err != nil {
				return err
			}
		} else if file.IsDir() {
			fileHeader.SetMode(file.Mode())
			_, err := w.CreateHeader(fileHeader)
			if err != nil {
				return err
			}
			err = addFiles(w, newPath, newInZip+"/")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

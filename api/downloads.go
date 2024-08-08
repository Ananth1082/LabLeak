package handler

import (
	"archive/zip"
	"net/http"
	"os"
	"path/filepath"
)

func sendDirContent(zw *zip.Writer, root string, basePath string) error {
	children, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, child := range children {
		filePath := filepath.Join(root, child.Name())
		zipPath := filepath.Join(basePath, child.Name())

		if child.IsDir() {
			// Recursively handle subdirectories
			err := sendDirContent(zw, filePath, zipPath)
			if err != nil {
				return err
			}
		} else {
			// Add files to the zip
			zipFile, err := zw.Create(zipPath)
			if err != nil {
				return err
			}
			fileBytes, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}
			_, err = zipFile.Write(fileBytes)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func DownloadScripts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=scripts.zip")
	zw := zip.NewWriter(w)
	defer zw.Close()

	err := sendDirContent(zw, "../scripts", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't send files: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DownloadExes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=executables.zip")
	zw := zip.NewWriter(w)
	defer zw.Close()

	err := sendDirContent(zw, "../executables", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't send files: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

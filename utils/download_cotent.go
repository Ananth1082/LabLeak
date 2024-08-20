package utils

import (
	"archive/zip"
	"os"
	"path/filepath"
)

func SendDirContentFromLocal(zw *zip.Writer, root string, basePath string) error {
	children, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, child := range children {
		filePath := filepath.Join(root, child.Name())
		zipPath := filepath.Join(basePath, child.Name())

		if child.IsDir() {
			// Recursively handle subdirectories
			err := SendDirContentFromLocal(zw, filePath, zipPath)
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

func SendDirContentFromDB(zw *zip.Writer, root map[string]string) error {
	for filePath, fileContent := range root {
		zipFile, err := zw.Create(filePath)
		if err != nil {
			return err
		}
		fileBytes := []byte(fileContent)
		_, err = zipFile.Write(fileBytes)
		if err != nil {
			return err
		}
	}
	return nil
}

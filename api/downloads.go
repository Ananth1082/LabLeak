package handler

import (
	"archive/zip"
	"net/http"

	"github.com/Ananth1082/LabLeak/utils"
)

func DownloadScripts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename=scripts.zip")
	zw := zip.NewWriter(w)
	defer zw.Close()

	err := utils.SendDirContentFromLocal(zw, "../scripts", "")
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

	err := utils.SendDirContentFromLocal(zw, "../executables", "")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't send files: " + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

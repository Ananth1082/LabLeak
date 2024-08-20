package handler

import (
	"fmt"
	"net/http"

	"github.com/Ananth1082/LabLeak/repository"
)

func CreateManual(w http.ResponseWriter, r *http.Request) {
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	manual := r.PathValue("manual")
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		http.Error(w, "file size exceeds 1Mb", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	content := make([]byte, header.Size)
	n, err := file.Read(content)
	fmt.Println("size read:", n)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details" + err.Error()))
		return
	}
	err = repository.CreateManual(section, subject, manual, header.Filename, string(content))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Succefully wrote the content"))
}

func DeleteManual(w http.ResponseWriter, r *http.Request) {
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	manual := r.PathValue("manual")

	err := repository.DeleteManual(section, subject, manual)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Succefully deleted the content"))
}

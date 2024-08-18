package handler

import (
	"fmt"
	"net/http"

	"github.com/Ananth1082/LabLeak/repository"
)

func GetManual(w http.ResponseWriter, r *http.Request) {
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	manual := r.PathValue("manual")
	manualContent, err := repository.GetManual(section, subject, manual)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details.."))
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(manualContent))
}

func CreateManual(w http.ResponseWriter, r *http.Request) {
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	manual := r.PathValue("manual")
	err := r.ParseMultipartForm(1 << 20)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
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
	err = repository.CreateManual(section, subject, manual, string(content))
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

func GetManuals(w http.ResponseWriter, r *http.Request) {
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	docs, err := repository.ListManuals(section, subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("*********List of manuals **********<br>"))
	for idx, doc := range docs {
		w.Write([]byte(fmt.Sprintf("%d. <a href=http://lableak.onrender.com/%s/%s/%s>%s</a><br>", idx+1, section, subject, doc.ID, doc.ID)))
	}
	w.Write([]byte("*********************************<br>"))
}

func GetSubjects(w http.ResponseWriter, r *http.Request) {
	section := r.PathValue("section")
	docs, err := repository.ListSubjects(section)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("**********List of subjects **********<br>"))
	for idx, doc := range docs {
		w.Write([]byte(fmt.Sprintf("%d. <a href=http://lableak.onrender.com/%s/%s>%s</a><br>", idx+1, section, doc.ID, doc.ID)))
	}
	w.Write([]byte("*********************************<br>"))
}

func GetSections(w http.ResponseWriter, r *http.Request) {
	docs, err := repository.ListSections()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("**********List of manuals **********<br>"))
	for idx, doc := range docs {
		w.Write([]byte(fmt.Sprintf("%d. <a href=http://lableak.onrender.com/%s>%s</a><br>", idx+1, doc.ID, doc.ID)))
	}
	w.Write([]byte("*********************************<br>"))
}

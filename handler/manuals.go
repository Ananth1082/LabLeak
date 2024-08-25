package handler

import (
	"fmt"
	"log"
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
	files := r.MultipartForm.File

	//read the code file
	fileHeaders := files["file"]
	if len(fileHeaders) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	fileHeader := fileHeaders[0]
	content := make([]byte, fileHeader.Size)
	file, err := fileHeader.Open()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	_, err = file.Read(content)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details" + err.Error()))
		return
	}

	//Read the attachements
	attachmentHeaders := files["attachments"]
	attachementIDs := make([]string, 0, len(attachmentHeaders))
	for _, attachmentHeader := range attachmentHeaders {
		atchContent := make([]byte, attachmentHeader.Size)
		file, err := attachmentHeader.Open()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid details"))
			return
		}
		n, err := file.Read(atchContent)
		fmt.Println("size read:", n)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid details" + err.Error()))
			return
		}
		attachemntID, err := repository.SendFile(atchContent, attachmentHeader.Filename)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error uploading images" + err.Error()))
			return
		}
		attachementIDs = append(attachementIDs, attachemntID)
	}
	log.Println(attachementIDs, len(attachementIDs))
	err = repository.CreateManual(section, subject, manual, fileHeader.Filename, string(content), attachementIDs)
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

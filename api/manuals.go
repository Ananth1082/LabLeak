package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Ananth1082/Lab_Manual/config"
)

func GetManual(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	manual := r.PathValue("manual")
	docSnap, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Get(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	data, _ := docSnap.DataAt("content")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(data.(string)))
}

func CreateManual(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
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
	filename := header.Filename
	fmt.Println("filename:", filename)
	if filename[len(filename)-3:] != "txt" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid file"))
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
	_, err = config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Create(ctx, map[string]string{"content": string(content)})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details" + err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Succefully wrote the content"))
}

func DeleteManual(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	manual := r.PathValue("manual")
	_, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Delete(ctx)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Succefully deleted the content"))
}

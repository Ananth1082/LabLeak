package handler

import (
	"context"
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
	file := r.FormValue("text")
	_, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Create(ctx, map[string]string{"content": file})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
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

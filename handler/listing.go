package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Ananth1082/LabLeak/repository"
)

func GetManuals(w http.ResponseWriter, r *http.Request) {
	userType := r.Header.Get("user-agent")
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	docs, err := repository.ListManuals(section, subject)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	if strings.HasPrefix(userType, "curl") {
		w.Write([]byte("******** List of manuals ********\n"))
		for idx, doc := range docs {
			w.Write([]byte(fmt.Sprintf("%d. %s\n", idx+1, doc.ID)))
		}
		w.Write([]byte("*********************************\n"))
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("*********List of manuals **********<br>"))
		for idx, doc := range docs {
			w.Write([]byte(fmt.Sprintf("%d. <a href=http://lableak.onrender.com/%s/%s/%s>%s</a><br>", idx+1, section, subject, doc.ID, doc.ID)))
		}
		w.Write([]byte("*********************************<br>"))
	}
}
func GetSubjects(w http.ResponseWriter, r *http.Request) {
	userType := r.Header.Get("user-agent")
	section := r.PathValue("section")
	docs, err := repository.ListSubjects(section)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}
	if strings.HasPrefix(userType, "curl") {
		w.Write([]byte("******** List of subjects ********\n"))
		for idx, doc := range docs {
			w.Write([]byte(fmt.Sprintf("%d. %s\n", idx+1, doc.ID)))
		}
		w.Write([]byte("*********************************\n"))
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("********** List of subjects **********<br>"))
		for idx, doc := range docs {
			w.Write([]byte(fmt.Sprintf("%d. <a href=http://lableak.onrender.com/%s/%s>%s</a><br>", idx+1, section, doc.ID, doc.ID)))
		}
		w.Write([]byte("*********************************<br>"))
	}
}

func GetSections(w http.ResponseWriter, r *http.Request) {
	userType := r.Header.Get("user-agent")
	docs, err := repository.ListSections()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details"))
		return
	}

	if strings.HasPrefix(userType, "curl") {
		w.Write([]byte("******** List of sections ********\n"))
		for idx, doc := range docs {
			w.Write([]byte(fmt.Sprintf("%d. %s\n", idx+1, doc.ID)))
		}
		w.Write([]byte("*********************************\n"))
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte("********** List of sections **********<br>"))
		for idx, doc := range docs {
			w.Write([]byte(fmt.Sprintf("%d. <a href=http://lableak.onrender.com/%s>%s</a><br>", idx+1, doc.ID, doc.ID)))
		}
		w.Write([]byte("*********************************<br>"))
	}
}

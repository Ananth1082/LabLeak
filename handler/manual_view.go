package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Ananth1082/LabLeak/repository"
)

var HTMLString = ""

func init() {
	bt, err := os.ReadFile("../HTML/code_template_format_string.txt")
	if err != nil {
		log.Fatal(err)
		return
	}
	HTMLString = string(bt)
}

func GetManual(w http.ResponseWriter, r *http.Request) {
	section := r.PathValue("section")
	subject := r.PathValue("subject")
	manual := r.PathValue("manual")
	userType := r.Header.Get("user-agent")
	manualContent, err := repository.GetManual(section, subject, manual)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid details..\n"))
		return
	}

	if strings.HasPrefix(userType, "curl") {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(manualContent))
	} else {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, HTMLString, manual, manual, manualContent)
	}
}

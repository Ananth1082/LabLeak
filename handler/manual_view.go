package handler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Ananth1082/LabLeak/repository"
	"github.com/Ananth1082/LabLeak/utils"
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
	manualContent, fileName, attachments, err := repository.GetManual(section, subject, manual)
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
		imgs := ""
		for _, atch := range attachments {
			imgs += fmt.Sprintf(`<div class="grid-item"><img src="%s" alt="Image 1"></div>`, utils.ConvertByteToURL(atch.Blob))
		}
		_, ext := utils.GetNameAndExt(fileName)
		fmt.Fprintf(w, HTMLString, ext, manual, fileName, ext, manualContent, imgs)
	}
}

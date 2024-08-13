package main

import (
	"log"
	"net/http"

	handler "github.com/Ananth1082/LabLeak/api"
	middleware "github.com/Ananth1082/LabLeak/middleware"
)

func main() {
	router := http.NewServeMux()
	//User routes
	router.HandleFunc("GET /ping", handler.Ping)
	router.HandleFunc("GET /", handler.GetSections)
	router.HandleFunc("GET /{section}", handler.GetSubjects)
	router.HandleFunc("GET /{section}/{subject}", handler.GetManuals)
	router.HandleFunc("GET /{section}/{subject}/{manual}", handler.GetManual)

	//download routes
	router.HandleFunc("GET /downloads/scripts", handler.DownloadScripts)
	router.HandleFunc("GET /downloads/executables", handler.DownloadExes)

	//protected routes
	router.HandleFunc("POST /{section}/{subject}/{manual}", handler.CreateManual)
	router.HandleFunc("DELETE /{section}/{subject}/{manual}", handler.DeleteManual)

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: middleware.CheckAdmin(middleware.Logging(router)),
	}
	log.Println("Server started in port: ", 8080)
	err := server.ListenAndServe()
	panic("Error: " + err.Error())
}

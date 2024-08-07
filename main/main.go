package main

import (
	"log"
	"net/http"

	handler "github.com/Ananth1082/Lab_Manual/api"
	middleware "github.com/Ananth1082/Lab_Manual/middleware"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", handler.Home)
	router.HandleFunc("GET /{section}/{subject}/{manual}", handler.GetManual)
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

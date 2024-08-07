package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		log.Println("ping")
		_, err := http.Get("https://studious-doodle.onrender.com")
		if err != nil {
			log.Println("error")
		} else {
			log.Println("pong")
		}
		time.Sleep(5 * time.Minute)
	}
}

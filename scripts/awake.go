package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	for {
		resTime := time.Now()
		log.Println("ping")
		_, err := http.Get("https://studious-doodle.onrender.com")
		if err != nil {
			log.Println("error")
		} else {
			log.Println("pong")
		}
		log.Println("Response time for restart: ", time.Since(resTime))
		log.Println("see you 5 mins later")

		time.Sleep(5 * time.Minute)
	}
}

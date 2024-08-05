package middleware

import (
	"log"
	"net/http"
	"time"
)

type specialWriter struct {
	http.ResponseWriter
	Statuscode int
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			sw := &specialWriter{
				ResponseWriter: w,
				Statuscode:     http.StatusOK,
			}

			next.ServeHTTP(sw, r)
			log.Println("Request from: ", r.RemoteAddr, " Method :", r.Method, " Time :", time.Since(start), "Path :", r.URL.Path, "Code :", sw.Statuscode)
		})
}

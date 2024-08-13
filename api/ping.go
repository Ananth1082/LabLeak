package handler

import "net/http"

func Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

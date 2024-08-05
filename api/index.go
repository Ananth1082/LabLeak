package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Ananth1082/Lab_Manual/config"
)

func Home(w http.ResponseWriter, r *http.Request) {
	docSnap, err := config.Firebase.Fs.Collection("test").Doc("foo").Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "<h1>Hello from Go!</h1>%v", docSnap.Data())
}

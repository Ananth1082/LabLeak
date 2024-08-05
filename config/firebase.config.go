package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

type firebaseClient struct {
	app *firebase.App
}

var fb *firebaseClient = new(firebaseClient)

func init() {
	var err error
	opt := option.WithCredentialsFile("../env/lab-manual-9dcc3-firebase-adminsdk-novaw-4b87b72bc0.json")
	fb.app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Error initializing app")
	}

}

func main() {
	fmt.Println(fb.app)
}

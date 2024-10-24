package config

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

type firebaseClient struct {
	app *firebase.App
	Fs  *firestore.Client
}

var Firebase *firebaseClient = new(firebaseClient)

func init() {
	var err error
	opt := option.WithCredentialsFile("../main/firebase.config.json")
	Firebase.app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		panic("Error initializing app")
	}
	Firebase.Fs, err = Firebase.app.Firestore(context.Background())
	if err != nil {
		panic("Error creating firestore instance," + err.Error())
	}
}

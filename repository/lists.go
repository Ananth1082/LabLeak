package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Ananth1082/LabLeak/config"
)

func ListSections() ([]*firestore.DocumentRef, error) {
	ctx := context.Background()
	docs, err := config.Firebase.Fs.Collection("sections").DocumentRefs(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	return docs, nil
}

func ListSubjects(section string) ([]*firestore.DocumentRef, error) {
	ctx := context.Background()
	docs, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").DocumentRefs(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	return docs, nil
}

func ListManuals(section, subject string) ([]*firestore.DocumentRef, error) {
	ctx := context.Background()
	docs, err := config.Firebase.Fs.Collection("sections").DocumentRefs(ctx).GetAll()
	if err != nil {
		return nil, err
	}
	return docs, nil
}

package repository

import (
	"context"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/Ananth1082/LabLeak/config"
)

func GetManual(section, subject, manual string) (string, string, error) {
	ctx := context.Background()
	docSnap, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Get(ctx)
	if err != nil {
		return "", "", err
	}
	data := docSnap.Data()

	manualContent := data["content"].(string)
	name, ok := data["name"].(string)
	if !ok {
		name = "code.txt"
	}

	return manualContent, name, nil
}

func CreateManual(section, subject, manual, fileName, content string) error {
	content = strings.Trim(content, "\n") + "\n"
	ctx := context.Background()
	_, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Create(ctx, map[string]string{"content": content, "name": fileName})
	if err != nil {
		return err
	}
	return nil
}

func DeleteManual(section, subject, manual string) error {
	ctx := context.Background()
	_, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSubject(section, subject string) error {
	ctx := context.Background()
	_, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteSection(section string) error {
	ctx := context.Background()
	_, err := config.Firebase.Fs.Collection("sections").Doc(section).Delete(ctx)
	if err != nil {
		return err
	}
	return nil
}

func UpdateManual(section, subject, manual, newContent string) error {
	ctx := context.Background()
	_, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Update(ctx, []firestore.Update{{Path: "content", Value: map[string]string{"content": newContent}}})
	return err
}

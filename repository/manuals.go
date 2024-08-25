package repository

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/Ananth1082/LabLeak/config"
)

func GetManual(section, subject, manual string) (string, string, []File, error) {
	ctx := context.Background()
	docSnap, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Get(ctx)
	if err != nil {
		return "", "", nil, err
	}
	data := docSnap.Data()

	manualContent := data["content"].(string)
	name, ok := data["name"].(string)
	if !ok {
		name = "code.txt"
	}
	ids, _ := data["attachments"].([]interface{})
	attachements := make([]File, 0, len(ids))
	for _, id := range ids {
		idstr := id.(string)
		file, err := GetFile(idstr)
		if err != nil {
			fmt.Println("Error accessing image, ID: ", idstr)
			continue
		}
		attachements = append(attachements, *file)
	}
	return manualContent, name, attachements, nil
}

func CreateManual(section, subject, manual, fileName, content string, attachmentIDs []string) error {
	content = strings.Trim(content, "\n") + "\n"
	ctx := context.Background()
	_, err := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual).Create(ctx, map[string]interface{}{"content": content, "name": fileName, "attachments": attachmentIDs})
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

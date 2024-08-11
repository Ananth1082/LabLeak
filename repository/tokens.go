package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/Ananth1082/LabLeak/config"
)

func CreateTokensForManual(section, subject, manual, password string) error {
	ctx := context.Background()
	doc := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return err
	}
	data, err := docSnap.DataAt("admin_tokens")
	if err != nil {
		_, err := doc.Update(ctx, []firestore.Update{{Path: "admin_token", Value: []string{password}}})
		if err != nil {
			return err
		}
	}
	adminTokens, _ := data.([]interface{})
	adminTokens = append(adminTokens, password)
	_, err = doc.Update(ctx, []firestore.Update{{Path: "admin_token", Value: adminTokens}})
	if err != nil {
		return err
	}
	return nil
}

func CreateTokensForSubject(section, subject, password string) error {
	ctx := context.Background()
	doc := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return err
	}
	data, err := docSnap.DataAt("admin_tokens")
	if err != nil {
		_, err := doc.Update(ctx, []firestore.Update{{Path: "admin_token", Value: []string{password}}})
		if err != nil {
			return err
		}
	}
	adminTokens, _ := data.([]interface{})
	adminTokens = append(adminTokens, password)
	_, err = doc.Update(ctx, []firestore.Update{{Path: "admin_token", Value: adminTokens}})
	if err != nil {
		return err
	}
	return nil
}
func CreateTokensForSection(section, password string) error {
	ctx := context.Background()
	doc := config.Firebase.Fs.Collection("sections").Doc(section)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return err
	}
	data, err := docSnap.DataAt("admin_tokens")
	if err != nil {
		_, err := doc.Update(ctx, []firestore.Update{{Path: "admin_token", Value: []string{password}}})
		if err != nil {
			return err
		}
	}
	adminTokens, _ := data.([]interface{})
	adminTokens = append(adminTokens, password)
	_, err = doc.Update(ctx, []firestore.Update{{Path: "admin_token", Value: adminTokens}})
	if err != nil {
		return err
	}
	return nil
}

func CheckTokensForManual(section, subject, manual, password string) (bool, error) {
	ctx := context.Background()
	doc := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject).Collection("manuals").Doc(manual)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return false, err
	}
	data, err := docSnap.DataAt("admin_tokens")
	if err != nil {
		return true, nil
	}
	adminTokens, _ := data.([]interface{})
	for _, tkn := range adminTokens {
		if password == tkn {
			return true, nil
		}
	}
	return false, nil
}
func CheckTokensForSubject(section, subject, password string) (bool, error) {
	ctx := context.Background()
	doc := config.Firebase.Fs.Collection("sections").Doc(section).Collection("subjects").Doc(subject)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return false, err
	}
	data, err := docSnap.DataAt("admin_tokens")
	if err != nil {
		return true, nil
	}
	adminTokens, _ := data.([]interface{})
	for _, tkn := range adminTokens {
		if password == tkn {
			return true, nil
		}
	}
	return false, nil
}
func CheckTokensForSection(section, password string) (bool, error) {
	ctx := context.Background()
	doc := config.Firebase.Fs.Collection("sections").Doc(section)
	docSnap, err := doc.Get(ctx)
	if err != nil {
		return false, err
	}
	data, err := docSnap.DataAt("admin_tokens")
	if err != nil {
		return true, nil
	}
	adminTokens, _ := data.([]interface{})
	for _, tkn := range adminTokens {
		if password == tkn {
			return true, nil
		}
	}
	return false, nil
}

package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Ananth1082/LabLeak/repository"
)

func UploadPhotos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Upload request received")

	// Get the image file from the form
	img, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Invalid image file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer img.Close() // Ensure file is closed

	// Allocate a byte slice of the correct size to hold the file content
	fileBytes := make([]byte, header.Size)

	// Read the image file into the byte slice
	_, err = io.ReadFull(img, fileBytes)
	if err != nil {
		log.Println("Error reading image file: ", err)
		http.Error(w, "Error reading image file", http.StatusInternalServerError)
		return
	}

	// Send the file to be stored in the database
	err = repository.SendFile(fileBytes, header.Filename)
	if err != nil {
		log.Println("Error uploading image: ", err)
		http.Error(w, "Error uploading image", http.StatusInternalServerError)
		return
	}

	// Respond to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully uploaded %s, file size: %d bytes", header.Filename, header.Size)
}

func GetPhoto(w http.ResponseWriter, r *http.Request) {
	photoID := r.PathValue("photoID")
	photo, err := repository.GetFile(photoID)
	if err != nil {
		log.Println("Error getting image: ", err)
		http.Error(w, "Error getting image", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "image/"+photo.Ext)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.%s\"", photo.Filename, photo.Ext))
	w.Header().Set("Cache-Control", "public, max-age=3600") // Cache for 1 hour
	w.WriteHeader(http.StatusOK)
	w.Write(photo.Blob)
}

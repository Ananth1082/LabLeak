package repository

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/Ananth1082/LabLeak/config"
	"github.com/Ananth1082/LabLeak/utils"
)

const PACKET_SIZE = 900 << 10 // 900 KB
type file struct {
	Filename string `firestore:"name"`
	Size     int64  `firestore:"size"`
	Ext      string `firestore:"ext"`
	Blob     []byte
}

func SendFile(fileContent []byte, fileName string) (string, error) {
	size := int64(len(fileContent))
	doc, err := initPhoto(fileName, size)
	if err != nil {
		return "", err
	}

	packetsCount := int(size / PACKET_SIZE)
	if size%PACKET_SIZE != 0 {
		packetsCount++
	}
	for i := 0; i < packetsCount; i++ {
		start := int64(i * PACKET_SIZE)
		end := start + PACKET_SIZE
		if end > size {
			end = size // Ensure we don't go out of bounds for the last packet
		}
		packet := fileContent[start:end]
		err := sendPacket(doc, packet, i+1)
		if err != nil {
			return "", err
		}
	}
	return doc.ID, nil
}

func GetFile(fileID string) (*file, error) {
	ctx := context.Background()

	// Reference to the main document
	doc := config.Firebase.Fs.Collection("photos").Doc(fileID)
	data, err := doc.Get(ctx)
	if err != nil {
		return nil, err
	}
	photo := new(file)
	err = data.DataTo(photo)
	if err != nil {
		return nil, err
	}
	// Prepare a slice to store the reconstructed file
	photo.Blob = make([]byte, 0, photo.Size)

	// Iterate through the content collection and reconstruct the file
	docItr := doc.Collection("content").Documents(ctx)
	for {
		docSnap, err := docItr.Next()
		if err != nil {
			break // Exit when all documents are processed
		}

		// Extract packet data
		var packet map[string]interface{}
		err = docSnap.DataTo(&packet)
		if err != nil {
			return nil, err
		}

		// Access the _byteString field properly
		packetData, ok := packet["bin"].([]byte)
		if !ok {
			// If it's not in byte format, try converting the base64 string manually
			packetStr, ok := packet["bin"].(string)
			if !ok {
				return nil, fmt.Errorf("invalid bin data format")
			}

			// Decode base64 back to binary data
			packetData, err = base64.StdEncoding.DecodeString(packetStr)
			if err != nil {
				log.Println("Error decoding base64 packet: ", err)
				return nil, err
			}
		}

		// Append the decoded packet to the final Blob
		photo.Blob = append(photo.Blob, packetData...)
	}
	return photo, nil
}

func initPhoto(filename string, fileSize int64) (*firestore.DocumentRef, error) {
	ctx := context.Background()
	newDoc := config.Firebase.Fs.Collection("photos").NewDoc()
	name, ext := utils.GetNameAndExt(filename)
	_, err := newDoc.Create(ctx, map[string]any{"name": name, "size": fileSize, "ext": ext})
	if err != nil {
		return nil, err
	}
	return newDoc, nil
}

func sendPacket(doc *firestore.DocumentRef, packet []byte, index int) error {
	ctx := context.Background()
	_, err := doc.Collection("content").Doc(fmt.Sprint(index)).Create(ctx, map[string][]byte{"bin": packet})
	if err != nil {
		return err
	}
	return nil
}

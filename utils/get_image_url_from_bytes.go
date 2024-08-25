package utils

import (
	"encoding/base64"
	"fmt"
)

func ConvertByteToURL(imgFile []byte) string {
	encodedImg := base64.StdEncoding.EncodeToString(imgFile)
	return fmt.Sprintf("data:image/jpeg;base64,%s", encodedImg)
}

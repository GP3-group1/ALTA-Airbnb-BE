package storage

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

const (
	LOCAL_BASE_URL string = "http://localhost:8083/"
	UPLOAD_PATH    string = "static/images/"
)

// UploadFile uploads an object
func UploadFile(file multipart.File, objectName string) (fileLocation string, err error) {
	// Destination
	dst, err := os.Create(filepath.Join(UPLOAD_PATH, filepath.Base(objectName)))
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}

	return LOCAL_BASE_URL + UPLOAD_PATH + objectName, nil
}

func DeleteFile(objectName string) error {
	err := os.Remove(strings.Replace(objectName, LOCAL_BASE_URL, "", 1))
	if err != nil {
		return err
	}

	return nil
}

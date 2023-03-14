package storage

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"

	"cloud.google.com/go/storage"
)

type ClientUploader struct {
	storageClient *storage.Client
	projectID     string
	bucketName    string
	uploadPath    string
}

var clientUploader *ClientUploader

func GetStorageClient() *ClientUploader{
	if clientUploader == nil {
		client, err := storage.NewClient(context.Background())
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}

		clientUploader = &ClientUploader{
			storageClient: client,
			bucketName:    "alta-airbnb",
			projectID:     "personal-374310",
			uploadPath:    "static/images/",
		}

		return clientUploader
	}
	return clientUploader
}

// UploadFile uploads an object
func (c *ClientUploader) UploadFile(file multipart.File, object string) (fileLocation string, err error) {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*300)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.storageClient.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	return wc.Name, nil
}

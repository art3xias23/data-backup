package google

import (
	"context"
	"fmt"
	"os"
	"time"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

type GoogleService struct {
	service *drive.Service
}

func createDriveService() (*GoogleService, error) {
	ctx := context.Background()
	client, err := google.DefaultClient(ctx, drive.DriveFileScope)
	if err != nil {
		return &GoogleService{}, fmt.Errorf("unable to create Google Drive API client: %v", err)
	}

	driveService, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return &GoogleService{}, fmt.Errorf("unable to create Google Drive API service: %v", err)
	}

	return &GoogleService{driveService}, nil
}

func (uploader *GoogleService) UploadFile(localFilePath, mimeType string) (string, error) {
	file, err := os.Open(localFilePath)
	if err != nil {
		return "", fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	driveFile := &drive.File{
		MimeType: mimeType,
		Name:     "UploadedFile.zip",
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uploadedFile, err := uploader.service.Files.Create(driveFile).Media(file).Do()
	if err != nil {
		return "", fmt.Errorf("Unable to upload file: %v", err)
	}

	return uploadedFile.Id, nil
}

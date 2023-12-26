package google

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const (
	ServiceAccount = "C:/Users/kmilchev/Desktop/Code/Go/go-backup/service.json" // Please set the json file of Service account.
	fileName       = "C:/Users/kmilchev/Desktop/Code/Go/go-backup/output.zip"
	parentFolderId = "1B7JUwn0YTNNfAXzKiMsVQdQmH03QKLgs"
	SCOPE          = drive.DriveScope
)

type GoogleService struct {
	service *drive.Service
}

func CreateDriveService() (*GoogleService, error) {
	ctx := context.Background()
	srv, err := drive.NewService(ctx, option.WithCredentialsFile(ServiceAccount), option.WithScopes(SCOPE))
	if err != nil {
		log.Fatalf("Warning: Unable to create drive Client %v", err)
	}

	return &GoogleService{srv}, nil
}

func (uploader *GoogleService) UploadFile(mimeType string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", fmt.Errorf("unable to open file: %v", err)
	}
	defer file.Close()

	info, _ := file.Stat()
	if err != nil {
		log.Fatalf("Warning: %v", err)
	}

	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	f := &drive.File{Name: info.Name(),
		Parents: []string{parentFolderId}}

	res, err := uploader.service.Files.
		Create(f).
		Media(file). //context.Background(), file, fileInf.Size(), baseMimeType).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()

	if err != nil {
		log.Fatalln(err)
	}

	return res.Name, nil
}

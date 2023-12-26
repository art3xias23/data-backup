package main

import (
	"fmt"
	"main/google"
	"main/lib"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// err := ZipFolder()

	// if err != nil {
	// 	return
	// }

	erra := UploadFolder()

	if erra != nil {
		return
	}
}

func ZipFolder() error {

	if err := godotenv.Load(); err != nil {
		lib.SendMail("Could not load env variables")
		return err
	}
	my_dir := os.Getenv("TinoBackUpPath")
	if my_dir == "" {
		err := fmt.Errorf("backUp path is empty, couldn ot proceed")
		lib.SendMail(err.Error())
		return err
	}
	// targetZipFile := "output.zip"
	start := time.Now()
	// destination, err := zipper.ZipIt(my_dir, my_dir, "output")
	err := lib.RecursiveZip(my_dir, "output.zip")
	if err != nil {
		fmt.Println("Error creating ZIP archive:", err)
		return err
	} else {
		fmt.Printf("ZIP archive created successfully.")
	}

	end := time.Now()

	duration := end.Sub(start)

	fmt.Printf("Time it took to zip: %v\n", duration)
	return nil
}

func UploadFolder() error {

	gService, err := google.CreateDriveService()

	if err != nil {
		fmt.Printf("Could not instantiate google service")
		return err
	}

	uploadedFileId, err := gService.UploadFile("application/zip")

	if err != nil {
		fmt.Printf("Could not upload file to drive")
		return err
	}

	fmt.Printf("Uploaded successful! File Id: %s", uploadedFileId)

	return nil
}

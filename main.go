package main

import (
	"fmt"
	"main/lib"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		lib.SendMail("Could not load env variables")
		return
	}
	my_dir := os.Getenv("TinoBackUpPath")
	if my_dir == "" {
		err := fmt.Errorf("backUp path is empty, couldn ot proceed")
		lib.SendMail(err.Error())
	}
	// targetZipFile := "output.zip"
	start := time.Now()
	err := lib.RecursiveZip(my_dir, "output.zip")
	if err != nil {
		fmt.Println("Error creating ZIP archive:", err)
	} else {
		fmt.Println("ZIP archive created successfully.")
	}

	end := time.Now()

	duration := end.Sub(start)

	fmt.Printf("Time it took to zip: %v\n", duration)
}

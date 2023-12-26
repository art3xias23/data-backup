package lib

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func RecursiveZip(source, target string) error {
	destinationFile, err := os.Create(target)
	if err != nil {
		return err
	}

	defer destinationFile.Close()

	writer := zip.NewWriter(destinationFile)
	defer writer.Close()
	return filepath.Walk(source,
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)

			if err != nil {
				return err
			}

			header.Method = zip.Deflate

			header.Name, err = filepath.Rel(filepath.Dir(source), path)

			if err != nil {
				return err
			}

			if info.IsDir() {
				header.Name += "/"
			}

			headerWriter, err := writer.CreateHeader(header)

			if err != nil {
				return err
			}

			if info.IsDir() {
				return err
			}

			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(headerWriter, f)
			return err
		})

}

func Zip2(origin string, destination string) error {

	destionationFile, _ := os.Create(destination)

	defer destionationFile.Close()

	writer := zip.NewWriter(destionationFile)

	defer writer.Close()

	return filepath.Walk(origin, func(path string, fi os.FileInfo, err error) error {
		header, err := zip.FileInfoHeader(fi)
	})
}

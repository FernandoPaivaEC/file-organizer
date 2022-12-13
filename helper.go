package main

import (
	"fmt"
	"os"
)

var months = [12]string{
	"Janeiro",
	"Fevereiro",
	"Março",
	"Abril",
	"Maio",
	"Junho",
	"Julho",
	"Agosto",
	"Setembro",
	"Outubro",
	"Novembro",
	"Dezembro",
}

func listFiles(dirPath string) ([]fileIndexType, error) {
	reader, err := os.Open(dirPath)

	if err != nil {
		return nil, err
	}

	files, err := reader.Readdir(0)

	if err != nil {
		return nil, err
	}

	var fileIndexArray []fileIndexType

	for _, file := range files {
		if !file.IsDir() {
			fileIndexItem := fileIndexType{
				name:  file.Name(),
				day:   fmt.Sprint(file.ModTime().Day()),
				month: months[file.ModTime().Month()-1],
				year:  fmt.Sprint(file.ModTime().Year()),
			}

			fileIndexArray = append(fileIndexArray, fileIndexItem)
		}
	}

	return fileIndexArray, nil
}

func createFolder(dirPath string) error {
	return os.MkdirAll("./"+dirPath, os.ModePerm)
}

func moveFile(sourcePath string, destinationPath string) error {
	return os.Rename("./"+sourcePath, "./"+destinationPath)
}

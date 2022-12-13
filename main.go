package main

import (
	"fmt"
	"os"
	"runtime"
)

type filesIndexType struct {
	name  string
	day   string
	month string
	year  string
}

func listFiles(dirPath string) ([]filesIndexType, error) {
	reader, err := os.Open(dirPath)

	if err != nil {
		return nil, err
	}

	files, err := reader.Readdir(0)

	if err != nil {
		return nil, err
	}

	var fileIndexArray []filesIndexType

	for _, file := range files {
		if !file.IsDir() {
			fileIndex := filesIndexType{
				name:  file.Name(),
				day:   fmt.Sprint(file.ModTime().Day()),
				month: file.ModTime().Month().String(),
				year:  fmt.Sprint(file.ModTime().Year()),
			}

			fileIndexArray = append(fileIndexArray, fileIndex)
		}
	}

	return fileIndexArray, nil
}

func createFolder(dirName string) error {
	return os.Mkdir("./"+dirName, 0777)
}

func moveFile(sourcePath string, destinationPath string) error {
	return os.Rename(sourcePath, destinationPath)
}

func getPathSeparator() string {
	pathSeparator := "/"

	if runtime.GOOS == "windows" {
		pathSeparator = "\\"
	}

	return pathSeparator
}

func main() {
	// listFiles("../../../../bin/../../mnt/c/Users/ferna")

	dirPath := os.Args[1]
	sortBy := os.Args[2]

	fileIndexArray, err := listFiles(dirPath)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if sortBy == "year" {
		for _, fileIndexItem := range fileIndexArray {
			createFolder(fileIndexItem.year)
		}
	}
}

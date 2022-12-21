package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

func listFiles(dirPath string) (FileIndex, error) {
	reader, err := os.Open(dirPath)

	if err != nil {
		return nil, err
	}

	files, err := reader.Readdir(0)

	if err != nil {
		return nil, err
	}

	var fileIndex FileIndex

	for _, file := range files {
		if !file.IsDir() {
			fileInfo := FileInfo{
				name: file.Name(),
				lastModified: LastModified{
					day:   fmt.Sprint(file.ModTime().Day()),
					month: months[file.ModTime().Month()-1],
					year:  fmt.Sprint(file.ModTime().Year()),
				},
			}

			fileIndex = append(fileIndex, fileInfo)
		}
	}

	return fileIndex, nil
}

func createFolder(dirPath string) error {
	return os.MkdirAll("./"+dirPath, os.ModePerm)
}

func moveFile(sourcePath string, destinationPath string) error {
	return os.Rename("./"+sourcePath, "./"+destinationPath)
}

func clearTerminal() {
	clearCommand := "clear"

	if runtime.GOOS == "windows" {
		clearCommand = "cls"
	}

	command := exec.Command(clearCommand)
	command.Stdout = os.Stdout
	command.Run()
}

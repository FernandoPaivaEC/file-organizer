package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	clearTerminal()

	fmt.Println(">>> Organizador de arquivos por data <<<")
	fmt.Println("v1.0")
	fmt.Println()

	if len(os.Args) < 3 {
		fmt.Println("Por favor, informe todos os argumentos:")
		fmt.Println()
		fmt.Println("Primeiro argumento -> Caminho da pasta a ser organizada")
		fmt.Println("Segundo argumento -> Nível de organização: ano | mês | dia")
		return
	}

	dirPath := os.Args[1]
	sortBy := os.Args[2]

	fileIndex, err := listFiles(dirPath)

	if err != nil {
		fmt.Println("Erro ao listar os arquivos:")
		fmt.Println(err.Error())
		return
	}

	if sortBy == "ano" {
		for _, fileInfo := range fileIndex {
			createFolder(fileInfo.lastModified.year)
			moveFile(
				fileInfo.name,
				filepath.Join(
					fileInfo.lastModified.year,
					fileInfo.name,
				),
			)
		}
	} else if sortBy == "mês" {
		for _, fileInfo := range fileIndex {
			createFolder(filepath.Join(
				fileInfo.lastModified.year,
				fileInfo.lastModified.month,
			))
			moveFile(
				fileInfo.name,
				filepath.Join(
					fileInfo.lastModified.year,
					fileInfo.lastModified.month,
					fileInfo.name,
				),
			)
		}
	} else if sortBy == "dia" {
		for _, fileInfo := range fileIndex {
			createFolder(filepath.Join(
				fileInfo.lastModified.year,
				fileInfo.lastModified.month,
				fileInfo.lastModified.day,
			))
			moveFile(
				fileInfo.name,
				filepath.Join(
					fileInfo.lastModified.year,
					fileInfo.lastModified.month,
					fileInfo.lastModified.day,
					fileInfo.name,
				),
			)
		}
	} else {
		fmt.Println("Por favor, informe uma das opções -> ano | mês | dia")
		return
	}
}

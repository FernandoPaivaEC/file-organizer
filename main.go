package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
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

	fileIndexArray, err := listFiles(dirPath)

	if err != nil {
		fmt.Println("Erro ao listar os arquivos:")
		fmt.Println(err.Error())
		return
	}

	if sortBy == "ano" {
		for _, fileIndexItem := range fileIndexArray {
			createFolder(fileIndexItem.year)
			moveFile(
				fileIndexItem.name,
				filepath.Join(
					fileIndexItem.year,
					fileIndexItem.name,
				),
			)
		}
	} else if sortBy == "mês" {
		for _, fileIndexItem := range fileIndexArray {
			createFolder(filepath.Join(
				fileIndexItem.year,
				fileIndexItem.month,
			))
			moveFile(
				fileIndexItem.name,
				filepath.Join(
					fileIndexItem.year,
					fileIndexItem.month,
					fileIndexItem.name,
				),
			)
		}
	} else if sortBy == "dia" {
		for _, fileIndexItem := range fileIndexArray {
			createFolder(filepath.Join(
				fileIndexItem.year,
				fileIndexItem.month,
				fileIndexItem.day,
			))
			moveFile(
				fileIndexItem.name,
				filepath.Join(
					fileIndexItem.year,
					fileIndexItem.month,
					fileIndexItem.day,
					fileIndexItem.name,
				),
			)
		}
	} else {
		fmt.Println("Por favor, informe uma das opções -> ano | mês | dia")
		return
	}
}

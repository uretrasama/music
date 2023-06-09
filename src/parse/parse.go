package parse

import (
	"os"
)

func ParseFileNames(dir []os.DirEntry) []string {
	fileNames := []string{}

	for _, file := range dir {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}

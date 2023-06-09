package parse

import (
	"../config"
	"fmt"
	"os"
)

func ParseFileNames(dir []os.DirEntry) []string {
	fileNames := []string{}

	for _, file := range dir {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}

func ParseToSlice() {
	// get argumetns and place to slice
	if len(os.Args) == 1 {
		config.Path = "."
	} else {
		config.Path = os.Args[1]
	}
	dir, errf := os.ReadDir(config.Path)
	if errf != nil {
		fmt.Println(errf)
		os.Exit(1)
	}
	config.FileNames = ParseFileNames(dir)
}

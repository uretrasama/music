package parse

import (
	"../config"
	"fmt"
	"os"
)

func ParseFileNames(dir []os.DirEntry) []string {
	fileNames := []string{}

	for _, file := range dir {
		if (string(rune(file.Name()[len(file.Name())-3])) + string(rune(file.Name()[len(file.Name())-2])) + string(rune(file.Name()[len(file.Name())-1]))) == "mp3" {
			fileNames = append(fileNames, file.Name())
		}
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

func MaxLenght(files []string) int {
	max := 0
	for _, file := range files {
		if len(file) > max {
			max = len(file)
		}
	}
	return max
}

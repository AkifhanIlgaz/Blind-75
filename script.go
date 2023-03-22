package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
)

var langExtensions = map[string]string{
	"Go":         ".go",
	"Javascript": ".js",
	"Python":     ".py",
	"Rust":       ".rs",
}

var (
	directory string
	file      string
)

func main() {

	flag.StringVar(&directory, "topic", "", "Name of the topic")
	flag.StringVar(&file, "file", "", "Name of file")
	flag.Parse()

	for lang, extension := range langExtensions {
		go createFile(lang, extension)
	}

}

func createFile(lang, extension string) {
	directoryPath := path.Join(lang, directory)
	fileName := file + extension

	err := os.MkdirAll(directoryPath, os.ModeDir)
	if err != nil {
		panic(err)
	}

	filePath := path.Join(directoryPath, fileName)

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {

		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		cmd := exec.Command("code", filePath)
		cmd.Run()

		defer file.Close()

	} else {
		fmt.Printf("%s already exists\n", filePath)
	}

}

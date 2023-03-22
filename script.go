package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"sync"
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
	wg        sync.WaitGroup
)

func main() {

	flag.StringVar(&directory, "topic", "", "Name of the topic")
	flag.StringVar(&file, "file", "", "Name of file")
	flag.Parse()

	wg.Add(len(langExtensions))

	for lang, extension := range langExtensions {
		go createFile(lang, extension)
	}
	
	wg.Wait()
}

func createFile(lang, extension string) {
	defer wg.Done()

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

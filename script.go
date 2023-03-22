package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

var langExtensions = map[string]string{
	"Go":         ".go",
	"Javascript": ".js",
	"Python":     ".py",
	"Rust":       ".rs",
}

func main() {
	var (
		directory string
		file      string
	)

	flag.StringVar(&directory, "topic", "", "Name of the topic")
	flag.StringVar(&file, "file", "", "Name of file")
	flag.Parse()

	for lang, extension := range langExtensions {
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
			defer file.Close()
		} else {
			fmt.Printf("%s already exists\n", filePath)
		}

	}

}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func findFile(dir string, name string) string {
	matchedPath := ""
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err, "-------------------")
		}

		if !f.IsDir() {
			if f.Name() == name {
				matchedPath = path
			}
		}
		return nil
	})
	if err != nil {
		log.Fatal(err, "-------------------")
	}
	if matchedPath == "" {
		log.Fatal(name, "does not exist")
	}

	return matchedPath
}

func main() {
	path := findFile("test", "file10")
	fmt.Println(path)
}

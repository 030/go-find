package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func find() string {
	return "hello"
}

func scanDir(name string) filepath.WalkFunc {
	return func(path string, f os.FileInfo, err error) error {
		// matchedPath := ""
		if err != nil {
			log.Fatal(err)
		}
		if !f.IsDir() {
			if f.Name() == name {
				fmt.Println(path)
				// matchedPath = path
			}
		}
		return nil
	}
}

func findFile(dir string, name string) {
	err := filepath.Walk(dir, scanDir(name))
	if err != nil {
		log.Fatal(err, "-------------------")
	}
}

func main() {
	findFile("test", "file0")
}

package main

import (
	"fmt"
	"log"
	"os"

	find "github.com/030/go-find/find"
)

func main() {
	if len(os.Args) <= 2 {
		log.Fatal("Usage: go-find <dir e.g. ~/.gradle> <filename e.g. commons-io-1.2.3>")
	}

	path, err := find.File(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
}

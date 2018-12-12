package find

import (
	"fmt"
	"os"
	"path/filepath"
)

// File is a function that is able return the path of a file that resides in a directory
func File(dir string, name string) (string, error) {
	matchedPath := ""
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.IsDir() {
			if f.Name() == name {
				matchedPath = path
			}
		}
		return nil
	})

	if matchedPath == "" {
		err = fmt.Errorf("File %s was not found in directory %s", name, dir)
	}

	return matchedPath, err
}

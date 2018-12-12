package find

import (
	"fmt"
	"os"
	"path/filepath"
)

func findFile(dir string, name string) (string, error) {
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

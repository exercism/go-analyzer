package assets

import (
	"net/http"

	"github.com/pkg/errors"
)

// GetDirs retrieces a list of sub directories of a given parent/path folder.
func GetDirs(path string, parent http.FileSystem) ([]string, error) {
	dir, err := parent.Open(path)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	files, err := dir.Readdir(0)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return dirs, nil
}

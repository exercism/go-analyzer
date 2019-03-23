package analyzer

import (
	"net/http"
	"path"

	"github.com/tehsphinx/astpatt"
	"github.com/tehsphinx/astrav"
)

// PatternDirs returns a list of folders with patterns for given exercise slug.
func PatternDirs(exercise string) ([]string, error) {
	paths, err := GetDirs(exercise, Patterns)
	if err != nil {
		return nil, err
	}

	for i, dir := range paths {
		paths[i] = path.Join("patterns", exercise, dir)
	}
	return paths, err
}

// CheckPattern checks if the given package matches any good pattern
func CheckPattern(exercise string, solution *astrav.Package) (float64, bool, error) {
	dirs, err := PatternDirs(exercise)
	if err != nil {
		return 0, false, err
	}
	patterns, err := loadPatterns(dirs...)
	_, ratio, ok := astpatt.DiffPatterns(patterns, solution)
	return ratio, ok, err
}

func loadPatterns(paths ...string) ([]*astpatt.Pattern, error) {
	var (
		err   error
		patts []*astpatt.Pattern
	)
	for _, dir := range paths {
		pkg, e := LoadPackage(dir)
		if e != nil {
			err = e
			continue
		}
		patts = append(patts, astpatt.ExtractPattern(pkg))
	}
	return patts, err
}

// GetDirs retrieces a list of sub directories of a given parent/path folder.
func GetDirs(path string, parent http.FileSystem) ([]string, error) {
	dir, err := parent.Open(path)
	if err != nil {
		return nil, err
	}

	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, file := range files {
		if file.IsDir() {
			dirs = append(dirs, file.Name())
		}
	}
	return dirs, nil
}

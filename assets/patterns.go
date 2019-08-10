package assets

import (
	"path"

	"github.com/pkg/errors"
	"github.com/tehsphinx/astpatt"
	"github.com/tehsphinx/astrav"
)

// LoadPatterns loads the patterns for given exercise.
func LoadPatterns(exercise string) ([]*astpatt.Pattern, error) {
	dirs, err := patternDirs(exercise)
	if err != nil {
		return nil, err
	}
	return loadPatterns(dirs...)
}

// patternDirs returns a list of folders with patterns for given exercise slug.
func patternDirs(exercise string) ([]string, error) {
	paths, err := GetDirs(exercise, Patterns)
	if err != nil {
		return nil, err
	}

	for i, dir := range paths {
		paths[i] = path.Join(exercise, dir)
	}
	return paths, err
}

func loadPatterns(paths ...string) ([]*astpatt.Pattern, error) {
	var (
		err   error
		patts []*astpatt.Pattern
	)
	for _, dir := range paths {
		pkg, e := loadPattern(dir)
		if e != nil {
			err = e
			continue
		}
		patts = append(patts, astpatt.ExtractPatternPermutations(pkg)...)
	}
	return patts, err
}

// loadPattern loads a go package from a folder
func loadPattern(path string) (*astrav.Package, error) {
	folder := astrav.NewFolder(Patterns, path)
	_, err := folder.ParseFolder()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return folder.Package(folder.Pkg.Name()), nil
}

package analyzer

import (
	"fmt"
	"net/http"
	"path"

	"github.com/exercism/go-analyzer/assets"
	"github.com/exercism/go-analyzer/suggester"
	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/pkg/errors"
	"github.com/tehsphinx/astpatt"
	"github.com/tehsphinx/astrav"
)

// Analyze analyses a solution and returns the Result
func Analyze(exercise string, path string) Result {
	var suggs = sugg.NewSuggestions()

	solution, err := LoadPackage(path)
	if err != nil {
		suggs.AppendUniquePH(sugg.SyntaxError, map[string]string{
			"err": fmt.Sprintf("%v", err),
		})
		return getResult(0, suggs)
	}
	if solution == nil {
		return NewErrResult(errors.New("there doesn't seem to be any solution uploaded"))
	}

	patternRating, _, err := CheckPattern(exercise, solution)
	if err != nil {
		return NewErrResult(err)
	}

	suggester.Suggest(exercise, solution, suggs)

	return getResult(patternRating, suggs)
}

// CheckPattern checks if the given package matches any good pattern
func CheckPattern(exercise string, solution *astrav.Package) (float64, bool, error) {
	patterns, err := assets.LoadPatterns(exercise)
	_, ratio, ok := astpatt.DiffPatterns(patterns, solution)
	return ratio, ok, err
}

// LoadPackage loads a go package from a folder
func LoadPackage(dir string) (*astrav.Package, error) {
	root := http.Dir(".")
	if path.IsAbs(dir) {
		root = http.Dir("/")
	}

	folder := astrav.NewFolder(root, dir)
	_, err := folder.ParseFolder()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return folder.Package(folder.Pkg.Name()), nil
}

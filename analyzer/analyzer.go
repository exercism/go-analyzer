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
		return getResult(exercise, nil, suggs)
	}
	if solution == nil {
		return NewErrResult(errors.New("there doesn't seem to be any solution uploaded"))
	}

	if exercise == "" {
		exercise = getExerciseSlug(solution)
	}

	pattReport, err := CheckPattern(exercise, solution)
	if err != nil {
		return NewErrResult(err)
	}

	suggester.Suggest(exercise, solution, suggs)

	return getResult(exercise, pattReport, suggs)
}

// PatternReport contain information on pattern matching
type PatternReport struct {
	// PatternRating is the pattern similarity with the closest pattern
	PatternRating float64
	// OptimalLimit is the exercise limit (similarity) for approve as optimal
	OptimalLimit float64
	// ApproveLimit is the exercise limit (similarity) for approve with comments
	ApproveLimit float64
	// Match is true if one of the patterns was a perfect match
	PerfectMatch bool
	// MinDiff is a git-like diff to the closest pattern
	MinDiff string
}

// CheckPattern checks if the given package matches any good pattern
func CheckPattern(exercise string, solution *astrav.Package) (*PatternReport, error) {
	patterns, err := assets.LoadPatterns(exercise)
	minDiff, ratio, ok := astpatt.DiffPatterns(patterns, solution)
	return &PatternReport{
		PatternRating: ratio,
		MinDiff:       minDiff,
		PerfectMatch:  ok,
	}, err
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

var pkgNameToSlug = map[string]string{
	"twofer": "two-fer",
}

func getExerciseSlug(pkg *astrav.Package) string {
	if slug, ok := pkgNameToSlug[pkg.Name]; ok {
		return slug
	}
	return pkg.Name
}

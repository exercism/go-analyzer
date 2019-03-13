package analyzer

import (
	"github.com/exercism/go-analyzer/assets"
	"github.com/tehsphinx/astpatt"
	"github.com/tehsphinx/astrav"
)

// CheckPattern checks if the given package matches any good pattern
func CheckPattern(exercise string, solution *astrav.Package) (bool, error) {
	dirs, err := assets.PatternDirs(exercise)
	if err != nil {
		return false, err
	}
	patterns := loadPatterns(dirs...)
	return astpatt.MatchPatterns(patterns, solution), nil
}

func loadPatterns(paths ...string) []*astpatt.Pattern {
	var patts []*astpatt.Pattern
	for _, path := range paths {
		pkg := LoadPackage(path)
		patts = append(patts, astpatt.ExtractPattern(pkg))
	}
	return patts
}

package analyzer

import (
	"errors"
	"fmt"

	"github.com/exercism/go-analyzer/suggester"
	"github.com/exercism/go-analyzer/suggester/sugg"
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

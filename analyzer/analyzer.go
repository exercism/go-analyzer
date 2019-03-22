package analyzer

import (
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
		return getResult(false, suggs)
	}

	goodPattern, err := CheckPattern(exercise, solution)
	if err != nil {
		return NewErrResult(err)
	}

	suggester.Suggest(exercise, solution, suggs)

	return getResult(goodPattern, suggs)
}

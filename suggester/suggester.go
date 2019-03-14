package suggester

import (
	"fmt"

	"github.com/exercism/go-analyzer/suggester/suggTypes"
	"github.com/exercism/go-analyzer/suggester/twofer"
	"github.com/tehsphinx/astrav"
)

var exercisePkgs = map[string][]suggTypes.SuggestionFunc{
	"twofer": twofer.FuncRegister,
}

// Suggest statically analysis the solution and returns a list of comments to provide.
func Suggest(exercise string, pkg *astrav.Package) (suggTypes.Suggestions, error) {
	funcs, ok := exercisePkgs[exercise]
	if !ok {
		return nil, fmt.Errorf("suggester for exercise '%s' not implemented", exercise)
	}

	var suggs suggTypes.Suggestions
	for _, fn := range funcs {
		res, err := fn(pkg)
		if err != nil {
			return nil, err
		}

		suggs.Merge(res)
	}
	return suggs, nil
}

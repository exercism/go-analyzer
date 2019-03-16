package suggester

import (
	"fmt"

	"github.com/exercism/go-analyzer/suggester/twofer"
	"github.com/exercism/go-analyzer/suggester/types"
	"github.com/tehsphinx/astrav"
)

var exercisePkgs = map[string]types.Register{
	"twofer": twofer.Register,
}

// Suggest statically analysis the solution and returns a list of comments to provide.
func Suggest(exercise string, pkg *astrav.Package) (*types.Suggestions, error) {
	register, ok := exercisePkgs[exercise]
	if !ok {
		return nil, fmt.Errorf("suggester for exercise '%s' not implemented", exercise)
	}

	var suggs = types.NewSuggestions(register.Severity)
	for _, fn := range register.Funcs {
		fn(pkg, suggs)
	}
	return suggs, nil
}

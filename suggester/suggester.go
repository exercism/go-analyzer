package suggester

import (
	"fmt"

	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/exercism/go-analyzer/suggester/twofer"
	"github.com/tehsphinx/astrav"
)

var exercisePkgs = map[string]sugg.Register{
	"twofer": twofer.Register,
}

// Suggest statically analysis the solution and returns a list of comments to provide.
func Suggest(exercise string, pkg *astrav.Package) *sugg.SuggestionReport {
	register, ok := exercisePkgs[exercise]
	if !ok {
		return nil
	}

	var suggs = sugg.NewSuggestions(register.Severity)
	for _, fn := range register.Funcs {
		func(fn sugg.SuggestionFunc) {
			defer func() {
				// in case one of the functions panics we catch that
				// and create an error from the panic value.
				if r := recover(); r != nil {
					suggs.ReportError(fmt.Errorf("PANIC: %+v", r))
				}
			}()

			fn(pkg, suggs)
		}(fn)
	}
	return suggs
}

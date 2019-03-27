package suggester

import (
	"fmt"
	"runtime/debug"

	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/exercism/go-analyzer/suggester/twofer"
	"github.com/tehsphinx/astrav"
)

var exercisePkgs = map[string]sugg.Register{
	"general": sugg.GeneralRegister,
	"two-fer": twofer.Register,
}

// Suggest statically analysis the solution and returns a list of comments to provide.
func Suggest(exercise string, pkg *astrav.Package, suggs *sugg.SuggestionReport) {
	suggs.SetSeverity(sugg.GeneralRegister.Severity)
	if pkg == nil {
		suggs.AppendUnique(sugg.SyntaxError)
		return
	}

	// TODO: compine severity general+exercise where exercise can overwrite general
	for _, key := range []string{"general", exercise} {
		register, ok := exercisePkgs[key]
		if !ok {
			continue
		}

		suggs.SetSeverity(register.Severity)
		for _, fn := range register.Funcs {
			catchSuggFunc(pkg, suggs, fn)
		}
	}
}

func catchSuggFunc(pkg *astrav.Package, suggs *sugg.SuggestionReport, fn sugg.SuggestionFunc) {
	defer func() {
		// in case one of the functions panics we catch that
		// and create an error from the panic value.
		if r := recover(); r != nil {
			suggs.ReportError(fmt.Errorf("PANIC: %+v\n%s", r, debug.Stack()))
		}
	}()

	fn(pkg, suggs)
}

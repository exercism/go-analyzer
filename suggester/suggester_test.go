package suggester

import (
	"errors"
	"testing"

	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/astrav"
)

var suggestTests = []struct {
	register sugg.Register
	errors   []error
}{
	{
		register: sugg.Register{
			Funcs: []sugg.SuggestionFunc{
				func(pkg *astrav.Package, suggs sugg.Suggester) {},
			},
		},
	},
	{
		register: sugg.Register{
			Funcs: []sugg.SuggestionFunc{
				func(pkg *astrav.Package, suggs sugg.Suggester) {
					panic(errors.New("some error"))
				},
			},
		},
		errors: []error{
			errors.New("PANIC: some error"),
		},
	},
	{
		register: sugg.Register{
			Funcs: []sugg.SuggestionFunc{
				func(pkg *astrav.Package, suggs sugg.Suggester) {
					panic(struct {
						test string
					}{
						test: "some object error",
					})
				},
			},
		},
		errors: []error{
			errors.New("PANIC: {test:some object error}"),
		},
	},
}

func TestSuggest(t *testing.T) {
	for _, test := range suggestTests {
		exercisePkgs = map[string]sugg.Register{
			"test": test.register,
		}

		suggs := sugg.NewSuggestions()
		Suggest("test", &astrav.Package{}, suggs)
		errs := suggs.GetErrors()
		assert.Equal(t, len(test.errors), len(errs))
		for i := 0; i < len(test.errors); i++ {
			assert.Contains(t, errs[i].Error(), test.errors[i].Error())
		}
	}
}

func TestSuggestUnknownPackage(t *testing.T) {
	suggs := sugg.NewSuggestions()
	Suggest("unknown", nil, suggs)
	errs := suggs.GetErrors()
	assert.Equal(t, []error(nil), errs)
}

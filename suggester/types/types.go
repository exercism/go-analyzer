package types

import (
	"github.com/tehsphinx/astrav"
)

// Register defines a register type to be provided by every suggerter track implementation.
// It contains the functions to be called to get
type Register struct {
	Funcs    []SuggestionFunc
	Severity map[string]int
}

// SuggestionFunc defines a function checking a solution for a specific problem.
type SuggestionFunc func(pkg *astrav.Package, suggs *Suggestions)

// Suggestion is a comment replied to a solution.
type Suggestion struct {
	Comment  string
	Severity int
}

// NewSuggestions creates a new collection of suggestions.
func NewSuggestions(severity map[string]int) *Suggestions {
	return &Suggestions{
		severity: severity,
	}
}

// Suggestions is a list of comments including severity information.
type Suggestions struct {
	suggs    []Suggestion
	severity map[string]int
}

// AppendUnique adds a comment. Can be called on a nil pointer and will not be added if it already exists.
func (s *Suggestions) AppendUnique(comment string) {
	s.appendUnique(comment)
}

// Merge merges another list of suggestions in. Duplicates will not be added.
func (s *Suggestions) Merge(suggs Suggestions) {
	for _, sugg := range suggs.suggs {
		s.appendUnique(sugg.Comment)
	}
}

func (s *Suggestions) appendUnique(comment string) {
	for _, sugg := range s.suggs {
		if sugg.Comment == comment {
			return
		}
	}

	s.suggs = append(s.suggs, Suggestion{
		Comment:  comment,
		Severity: s.severity[comment],
	})
}

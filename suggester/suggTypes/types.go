package suggTypes

import (
	"github.com/tehsphinx/astrav"
)

// SuggestionFunc defines a function checking a solution for a specific problem.
type SuggestionFunc func(pkg *astrav.Package) (Suggestions, error)

// Suggestion is a comment replied to a solution.
type Suggestion struct {
	Comment  string
	Severity int
}

// Suggestions is a list of comments including severity information.
type Suggestions []Suggestion

// AppendUnique adds a comment. Can be called on a nil pointer and will not be added if it already exists.
func (s *Suggestions) AppendUnique(comment string, severityDefinitions map[string]int) {
	s.appendUnique(comment, severityDefinitions[comment])
}

// Merge merges another list of suggestions in. Duplicates will not be added.
func (s *Suggestions) Merge(suggs Suggestions) {
	for _, sugg := range suggs {
		s.appendUnique(sugg.Comment, sugg.Severity)
	}
}

func (s *Suggestions) appendUnique(comment string, severity int) {
	for _, sugg := range *s {
		if sugg.Comment == comment {
			return
		}
	}

	*s = append(*s, Suggestion{
		Comment:  comment,
		Severity: severity,
	})
}

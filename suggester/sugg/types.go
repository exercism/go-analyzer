package sugg

import (
	"github.com/tehsphinx/astrav"
)

// Suggester defines a list of comments including severity information.
// The reason for this interface is mainly to provide a focused (limited)
// set of functionality to suggester implementers.
type Suggester interface {
	// AppendUnique adds a suggestion while checking if it exists already.
	// That way it does not matter if the code accidentally adds the same suggestion multiple times.
	AppendUnique(comment string)

	// ReportError collects provided errors. They will be added to the output file
	// for debugging purpose. Reporting will fail the analyzer with `refer_to_mentor` status.
	// If that is not what you want consider adding a comment to the student instead of an error.
	ReportError(err error)
}

// Register defines a register type to be provided by every suggerter track implementation.
// It contains the functions to be called to get
type Register struct {
	// Funcs a registry of functions to be called. Each function should investigate one pattern and
	// can add one or multiple suggestions if the found pattern needs commenting.
	Funcs []SuggestionFunc

	// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
	// The maximum for a single comment is 5. A comment with that severity will block approval.
	// When assigning the severity a good guideline is to ask: How many comments of similar severity
	// should block approval?
	// We can be very strict on automated comments since the student has a very fast feedback loop.
	Severity map[string]int
}

// SuggestionFunc defines a function checking a solution for a specific problem.
type SuggestionFunc func(pkg *astrav.Package, suggs Suggester)

type suggestion struct {
	comment  string
	severity int
}

// NewSuggestions creates a new collection of suggestions.
func NewSuggestions(severity map[string]int) *SuggestionReport {
	return &SuggestionReport{
		severity: severity,
	}
}

// SuggestionReport is a list of comments including severity information.
type SuggestionReport struct {
	suggs    []suggestion
	severity map[string]int
	errors   []error
}

// AppendUnique adds a comment if it does not exist.
func (s *SuggestionReport) AppendUnique(comment string) {
	s.appendUnique(comment)
}

// ReportError reports an error to the analyzer.
func (s *SuggestionReport) ReportError(err error) {
	s.errors = append(s.errors, err)
}

// GetComments returns the comments and their severity sum.
func (s *SuggestionReport) GetComments() ([]string, int) {
	if s == nil {
		return nil, 0
	}

	var (
		comments    []string
		sumSeverity int
	)
	for _, sugg := range s.suggs {
		comments = append(comments, sugg.comment)
		sumSeverity += sugg.severity
	}
	return comments, sumSeverity
}

// GetErrors returns a list of errors that occured.
func (s *SuggestionReport) GetErrors() []error {
	if s == nil {
		return nil
	}
	return s.errors
}

func (s *SuggestionReport) appendUnique(comment string) {
	for _, sugg := range s.suggs {
		if sugg.comment == comment {
			return
		}
	}

	s.suggs = append(s.suggs, suggestion{
		comment:  comment,
		severity: s.severity[comment],
	})
}

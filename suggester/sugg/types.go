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

	// AppendUniquePH adds a suggestion with placeholders while checking if it exists already.
	AppendUniquePH(commentID string, params map[string]string)

	// ReportError collects provided errors. They will be added to the output file
	// for debugging purpose. Reporting will fail the analyzer with `refer_to_mentor` status.
	// If that is not what you want consider adding a comment to the student instead of an error.
	ReportError(err error)

	// HasSuggestion checks if a comment was added. This should not be used to avoid duplicated.
	// Duplicates are avoided by default. It might however be useful to check if some other algorithm
	// found a certain pattern so it doesn't have to be checked again.
	HasSuggestion(comment string) bool
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

// NewSuggestions creates a new collection of suggestions.
func NewSuggestions() *SuggestionReport {
	return &SuggestionReport{
		severity: GeneralRegister.Severity,
	}
}

// SuggestionReport is a list of comments including severity information.
type SuggestionReport struct {
	suggs    []Comment
	severity map[string]int
	errors   []error
}

// AppendUnique adds a comment if it does not exist.
func (s *SuggestionReport) AppendUnique(commentID string) {
	s.appendUnique(NewComment(commentID))
}

// AppendUniquePH adds a comment with placeholder(s). Uniqueness includes the placeholder(s) and value(s).
func (s *SuggestionReport) AppendUniquePH(commentID string, params map[string]string) {
	s.appendUnique(NewPlaceholderComment(commentID, params))
}

// AppendBlock adds a block comment if it does not exist.
func (s *SuggestionReport) AppendBlock(commentID string) {
	s.appendUnique(NewBlockComment(commentID))
}

// ReportError reports an error to the analyzer.
func (s *SuggestionReport) ReportError(err error) {
	if err == nil {
		return
	}
	s.errors = append(s.errors, err)
}

// HasSuggestion checks if the comment was added already. Params are ignored for comparison.
func (s *SuggestionReport) HasSuggestion(commentID string) bool {
	for _, cmt := range s.suggs {
		if cmt.ID() == commentID {
			return true
		}
	}
	return false
}

// GetComments returns the comments and their severity sum.
func (s *SuggestionReport) GetComments() ([]Comment, int) {
	if s == nil {
		return nil, 0
	}

	var (
		comments    []Comment
		sumSeverity int
	)
	for _, sugg := range s.suggs {
		comments = append(comments, sugg)
		sumSeverity += sugg.Severity()
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

// AppendSeverity adds new severities overwriting existing ones.
func (s *SuggestionReport) AppendSeverity(severity map[string]int) {
	for k, v := range severity {
		s.severity[k] = v
	}
}

func (s *SuggestionReport) appendUnique(comment Comment) {
	if Contains(s.suggs, comment) {
		return
	}

	comment.setSeverity(s.severity[comment.ID()])
	s.suggs = append(s.suggs, comment)
}

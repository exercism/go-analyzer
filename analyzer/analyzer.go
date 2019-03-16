package analyzer

import (
	"github.com/exercism/go-analyzer/suggester"
	"github.com/exercism/go-analyzer/suggester/sugg"
)

// Result defines the result of the analyzer
type Result struct {
	Status   Status   `json:"status"`
	Comments []string `json:"comments"`
	Errors   []error  `json:"errors"`
}

// Status defines the status of a solution to be acted upon by exercism.
type Status string

// status constants
const (
	ApproveAsOptimal      Status = "approve_as_optimal"
	ApproveWithComment    Status = "approve_with_comment"
	DisapproveWithComment Status = "disapprove_with_comment"
	ReferToMentor         Status = "refer_to_mentor"
)

// Analyze analyses a solution and returns the Result
func Analyze(exercise string, path string) Result {
	solution, err := LoadPackage(path)
	if err != nil {
		return Result{Errors: []error{err}}
	}

	goodPattern, err := CheckPattern(exercise, solution)
	if err != nil {
		return Result{Errors: []error{err}}
	}

	suggs := suggester.Suggest(exercise, solution)

	return getResult(goodPattern, suggs)
}

func getResult(goodPattern bool, suggReporter *sugg.SuggestionReport) Result {
	comments, severity := suggReporter.GetComments()
	if comments == nil {
		// make sure not to add nil to json
		comments = []string{}
	}
	errs := suggReporter.GetErrors()
	return Result{
		Status:   getStatus(goodPattern, len(comments), severity, len(errs)),
		Comments: comments,
		Errors:   errs,
	}
}

func getStatus(goodPattern bool, comments, severity int, errors int) Status {
	switch {
	case errors != 0:
		return ReferToMentor
	case goodPattern && comments == 0:
		return ApproveAsOptimal
	case goodPattern && severity < 5:
		return ApproveWithComment
	case goodPattern:
		return DisapproveWithComment
	case 5 <= severity:
		return DisapproveWithComment
	}

	return ReferToMentor
}

package analyzer

import (
	"fmt"

	"github.com/exercism/go-analyzer/suggester/sugg"
)

// NewErrResult creates a result with an error.
func NewErrResult(err error) Result {
	return Result{
		Comments: []sugg.Comment{},
		Errors:   []string{fmt.Sprintf("%v", err)},
	}
}

// Result defines the result of the analyzer
type Result struct {
	Status   Status         `json:"status"`
	Comments []sugg.Comment `json:"comments"`
	Errors   []string       `json:"errors,omitempty"`
	Severity int            `json:"-"`
}

// Status defines the status of a solution to be acted upon by exercism.
type Status string

// status constants
const (
	ApproveAsOptimal      Status = "approve_as_optimal"
	ApproveWithComment    Status = "approve_with_comment"
	DisapproveWithComment Status = "disapprove_with_comment"
	ReferToMentor         Status = "refer_to_mentor"
	Ejected               Status = "ejected"
)

func getResult(goodPattern bool, suggReporter *sugg.SuggestionReport) Result {
	comments, severity := suggReporter.GetComments()
	if comments == nil {
		// make sure not to add nil to json
		comments = []sugg.Comment{}
	}
	errs := suggReporter.GetErrors()
	return Result{
		Status:   getStatus(goodPattern, len(comments), severity, len(errs)),
		Comments: comments,
		Errors:   fmtErrors(errs),
		Severity: severity,
	}
}

func fmtErrors(errs []error) []string {
	var strs []string
	for _, err := range errs {
		strs = append(strs, fmt.Sprintf("%s", err))
	}
	return strs
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

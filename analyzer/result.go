package analyzer

import (
	"fmt"

	"github.com/exercism/go-analyzer/suggester/sugg"
)

// NewErrResult creates a result with an error.
func NewErrResult(err error) Result {
	return Result{
		Status:   ReferToMentor,
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
	Rating   float64        `json:"-"`
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

func getResult(patternRating float64, suggReporter *sugg.SuggestionReport) Result {
	comments, severity := suggReporter.GetComments()
	if comments == nil {
		// make sure not to add nil to json
		comments = []sugg.Comment{}
	}
	errs := suggReporter.GetErrors()
	return Result{
		Status:   getStatus(patternRating, len(comments), severity, len(errs)),
		Comments: comments,
		Errors:   fmtErrors(errs),
		Severity: severity,
		Rating:   patternRating,
	}
}

func fmtErrors(errs []error) []string {
	var strs []string
	for _, err := range errs {
		strs = append(strs, fmt.Sprintf("%s", err))
	}
	return strs
}

var (
	limitGoodPattern    = 0.99
	limitAllowedPattern = 0.90
)

func getStatus(rating float64, comments, severity int, errors int) Status {
	switch {
	case errors != 0:
		return ReferToMentor

	case comments == 0 && rating <= limitGoodPattern:
		return ReferToMentor
	case comments == 0 && limitGoodPattern < rating:
		return ApproveAsOptimal

	case limitAllowedPattern < rating && severity < 5:
		return ApproveWithComment
	case limitAllowedPattern < rating:
		return DisapproveWithComment
	case 5 <= severity:
		return DisapproveWithComment
	}

	return ReferToMentor
}

package analyzer

import (
	"github.com/exercism/go-analyzer/suggester"
)

// Result defines the result of the analyzer
type Result struct {
	Status   Status   `json:"status"`
	Comments []string `json:"comments"`
	Errors   []error  `json:"errors,omitempty"`
	Severity int      `json:"-"`
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

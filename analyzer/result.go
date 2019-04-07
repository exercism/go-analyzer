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
		Errors:   []string{fmt.Sprintf("%+v", err)},
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

type patternReport struct {
	// PatternRating is the pattern similarity with the closest pattern
	PatternRating float64
	// OptimalLimit is the exercise limit (similarity) for approve as optimal
	OptimalLimit float64
	// ApproveLimit is the exercise limit (similarity) for approve with comments
	ApproveLimit float64
}

func getResult(exercise string, patternRating float64, suggReporter *sugg.SuggestionReport) Result {
	var (
		comments, severity = suggReporter.GetComments()
		errs               = suggReporter.GetErrors()
		report             = getPatternReport(exercise, patternRating)
	)
	if comments == nil {
		// make sure not to add nil to json
		comments = []sugg.Comment{}
	}
	if report.OptimalLimit == 0 {
		panic(fmt.Sprintf("Programming Error: missing pattern limits for `%s`", exercise))
	}

	return Result{
		Status:   getStatus(report, len(comments), severity, len(errs)),
		Comments: comments,
		Errors:   fmtErrors(errs),
		Severity: severity,
		Rating:   patternRating,
	}
}

func getPatternReport(exercise string, rating float64) patternReport {
	report := patternLimits[exercise]
	report.PatternRating = rating
	return report
}

func fmtErrors(errs []error) []string {
	var strs []string
	for _, err := range errs {
		strs = append(strs, fmt.Sprintf("%+v", err))
	}
	return strs
}

func getStatus(pattern patternReport, comments, severity int, errors int) Status {
	switch {
	case errors != 0:
		// Some error(s) occured. Better leave it to a mentor.
		return ReferToMentor

	case comments == 0 && pattern.OptimalLimit < pattern.PatternRating:
		// The code is close enough to be approved as optimal and we have no suggestions.
		return ApproveAsOptimal
	case comments == 0 && pattern.PatternRating <= pattern.OptimalLimit:
		// The code is not close enough to be approved as optimal, but we don't know how to improve it.
		return ReferToMentor

	case pattern.ApproveLimit < pattern.PatternRating && severity < 5:
		// The code is close enough to approve, but we have minor improvement suggestions.
		return ApproveWithComment
	case pattern.ApproveLimit < pattern.PatternRating:
		// The code is close to a good solution, but we found too much or bigger things to improve on.
		return DisapproveWithComment
	case 5 <= severity:
		// The code is not close to a good solution, but the analyzer has enough suggestions to improve on.
		return DisapproveWithComment
	}

	// Default: Better leave it to a mentor.
	return ReferToMentor
}

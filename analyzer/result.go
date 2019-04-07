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
	PatternRating float64
	OptimalLimit  float64
	ApproveLimit  float64
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
		return ReferToMentor

	case comments == 0 && pattern.PatternRating <= pattern.OptimalLimit:
		return ReferToMentor
	case comments == 0 && pattern.OptimalLimit < pattern.PatternRating:
		return ApproveAsOptimal

	case pattern.ApproveLimit < pattern.PatternRating && severity < 5:
		return ApproveWithComment
	case pattern.ApproveLimit < pattern.PatternRating:
		return DisapproveWithComment
	case 5 <= severity:
		return DisapproveWithComment
	}

	return ReferToMentor
}

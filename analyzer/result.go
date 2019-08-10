package analyzer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

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
	MinDiff  string         `json:"-"`
}

// Status defines the status of a solution to be acted upon by exercism.
type Status string

// status constants
const (
	Approve       Status = "approve"
	Disapprove    Status = "disapprove"
	ReferToMentor Status = "refer_to_mentor"
	Ejected       Status = "ejected"
)

func getResult(exercise string, pattReport *PatternReport, suggReporter *sugg.SuggestionReport) Result {
	var (
		comments, severity = suggReporter.GetComments()
		errs               = suggReporter.GetErrors()
	)
	if pattReport == nil {
		pattReport = &PatternReport{}
	}
	addLimits(exercise, pattReport)
	if comments == nil {
		// make sure not to add nil to json
		comments = []sugg.Comment{}
	}
	if pattReport.OptimalLimit == 0 {
		panic(fmt.Sprintf("Programming Error: missing pattern limits for `%s`", exercise))
	}

	return Result{
		Status:   getStatus(pattReport, len(comments), severity, len(errs)),
		Comments: comments,
		Errors:   fmtErrors(errs),
		Severity: severity,
		Rating:   pattReport.PatternRating,
		MinDiff:  pattReport.MinDiff,
	}
}

func addLimits(exercise string, report *PatternReport) *PatternReport {
	limits := patternLimits[exercise]
	report.OptimalLimit = limits.OptimalLimit
	report.ApproveLimit = limits.ApproveLimit
	return report
}

func fmtErrors(errs []error) []string {
	var strs []string
	for _, err := range errs {
		strs = append(strs, fmt.Sprintf("%+v", err))
	}
	return strs
}

func getStatus(pattern *PatternReport, comments, severity int, errors int) Status {
	switch {
	case errors != 0:
		// Some error(s) occured. Better leave it to a mentor.
		return ReferToMentor
	case comments == 0 && pattern.PatternRating <= pattern.OptimalLimit:
		// The code is not close enough to be approved as optimal, but we don't know how to improve it.
		return ReferToMentor

	case pattern.ApproveLimit < pattern.PatternRating && severity < 5:
		// The code is good enough to approve and we have no or minor improvement suggestions.
		return Approve
	case pattern.ApproveLimit < pattern.PatternRating:
		// The code is close to a good solution, but we found too much or bigger things to improve on.
		return Disapprove
	case 5 <= severity:
		// The code is not close to a good solution, but the analyzer has enough suggestions to improve on.
		return Disapprove
	}

	// Default: Better leave it to a mentor.
	return ReferToMentor
}

// GetResultFromFile returns the content of the `expected.json` file in given path.
func GetResultFromFile(dir string) (*Result, error) {
	bytes, err := ioutil.ReadFile(path.Join(dir, "expected.json"))
	if err != nil {
		return nil, err
	}

	// transforming to struct and back to json to eliminate different formatting
	var res = unmarshalResult{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		return nil, err
	}

	result := &Result{
		Status:   res.Status,
		Severity: res.Severity,
		Errors:   res.Errors,
	}
	for _, comment := range res.Comments {
		switch cmt := comment.(type) {
		case string:
			result.Comments = append(result.Comments, sugg.NewComment(cmt))
		case map[string]interface{}:
			comment, _ := cmt["comment"].(string)
			ps, _ := cmt["params"].(map[string]interface{})

			params := map[string]string{}
			for key, value := range ps {
				params[key], _ = value.(string)
			}

			result.Comments = append(result.Comments, sugg.NewPlaceholderComment(comment, params))
		}
	}
	return result, err
}

type unmarshalResult struct {
	Status   Status        `json:"status"`
	Comments []interface{} `json:"comments"`
	Errors   []string      `json:"errors,omitempty"`
	Severity int           `json:"-"`
}

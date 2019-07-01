package analyzer

import (
	"errors"
	"testing"

	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/stretchr/testify/assert"
)

var getResultTests = []struct {
	exercise    string
	goodPattern *PatternReport
	comments    []string
	result      Result
	severity    map[string]int
	errors      []error
}{
	{
		exercise:    "two-fer",
		goodPattern: &PatternReport{PatternRating: 1},
		result: Result{
			Status:   ApproveAsOptimal,
			Comments: []sugg.Comment{},
		},
	},
	{
		exercise:    "two-fer",
		goodPattern: &PatternReport{PatternRating: 0},
		result: Result{
			Status:   ReferToMentor,
			Comments: []sugg.Comment{},
		},
	},
	{
		exercise:    "two-fer",
		goodPattern: &PatternReport{PatternRating: 0},
		comments:    []string{"go.two-fer.some_comment"},
		result: Result{
			Status:   ReferToMentor,
			Comments: []sugg.Comment{sugg.NewComment("go.two-fer.some_comment")},
		},
	},
	{
		exercise:    "two-fer",
		goodPattern: &PatternReport{PatternRating: 1},
		comments:    []string{"go.two-fer.some_comment"},
		result: Result{
			Status:   ApproveWithComment,
			Comments: []sugg.Comment{sugg.NewComment("go.two-fer.some_comment")},
		},
	},
	{
		exercise:    "two-fer",
		goodPattern: &PatternReport{PatternRating: 1},
		comments:    []string{"go.two-fer.some_comment"},
		result: Result{
			Status:   DisapproveWithComment,
			Comments: []sugg.Comment{sugg.NewComment("go.two-fer.some_comment")},
			Severity: 5,
		},
		severity: map[string]int{"go.two-fer.some_comment": 5},
	},
	{
		exercise:    "two-fer",
		goodPattern: &PatternReport{PatternRating: 1},
		result: Result{
			Status:   ReferToMentor,
			Comments: []sugg.Comment{},
			Errors:   []string{"some error"},
		},
		errors: []error{errors.New("some error")},
	},
	{
		exercise:    "two-fer",
		goodPattern: &PatternReport{PatternRating: 1},
		result: Result{
			Status:   ApproveAsOptimal,
			Comments: []sugg.Comment{},
		},
		errors: []error{nil},
	},
	{
		exercise:    "two-fer",
		goodPattern: &PatternReport{PatternRating: 0},
		comments: []string{
			"go.two-fer.some_comment",
			"go.two-fer.some_comment_2",
			"go.two-fer.some_comment_3",
		},
		result: Result{
			Status: DisapproveWithComment,
			Comments: []sugg.Comment{
				sugg.NewComment("go.two-fer.some_comment"),
				sugg.NewComment("go.two-fer.some_comment_2"),
				sugg.NewComment("go.two-fer.some_comment_3"),
			},
			Severity: 6,
		},
		severity: map[string]int{
			"go.two-fer.some_comment":   2,
			"go.two-fer.some_comment_2": 1,
			"go.two-fer.some_comment_3": 3,
		},
	},
}

func Test_getResult(t *testing.T) {
	for _, test := range getResultTests {
		suggs := sugg.NewSuggestions()
		suggs.AppendSeverity(test.severity)
		for _, comment := range test.comments {
			suggs.AppendUnique(comment)
		}
		for _, err := range test.errors {
			suggs.ReportError(err)
		}

		res := getResult(test.exercise, test.goodPattern, suggs)

		assert.Equal(t, test.result.Status, res.Status)
		assert.Equal(t, test.result.Severity, res.Severity)
		assert.Equal(t, test.result.Errors, res.Errors)

		assert.Equal(t, len(test.result.Comments), len(res.Comments))
		for _, comment := range test.result.Comments {
			assert.True(t, sugg.Contains(res.Comments, comment))
		}
	}
}

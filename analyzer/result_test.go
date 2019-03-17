package analyzer

import (
	"errors"
	"testing"

	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/stretchr/testify/assert"
)

var getResultTests = []struct {
	goodPattern bool
	comments    []string
	result      Result
	severity    map[string]int
	errors      []error
}{
	{
		goodPattern: true,
		result: Result{
			Status:   ApproveAsOptimal,
			Comments: []string{},
		},
	},
	{
		goodPattern: false,
		result: Result{
			Status:   ReferToMentor,
			Comments: []string{},
		},
	},
	{
		goodPattern: false,
		comments:    []string{"go.two-fer.some_comment"},
		result: Result{
			Status:   ReferToMentor,
			Comments: []string{"go.two-fer.some_comment"},
		},
	},
	{
		goodPattern: true,
		comments:    []string{"go.two-fer.some_comment"},
		result: Result{
			Status:   ApproveWithComment,
			Comments: []string{"go.two-fer.some_comment"},
		},
	},
	{
		goodPattern: true,
		comments:    []string{"go.two-fer.some_comment"},
		result: Result{
			Status:   DisapproveWithComment,
			Comments: []string{"go.two-fer.some_comment"},
			Severity: 5,
		},
		severity: map[string]int{"go.two-fer.some_comment": 5},
	},
	{
		goodPattern: true,
		result: Result{
			Status:   ReferToMentor,
			Comments: []string{},
			Errors:   []error{errors.New("some error")},
		},
		errors: []error{errors.New("some error")},
	},
	{
		goodPattern: true,
		result: Result{
			Status:   ApproveAsOptimal,
			Comments: []string{},
		},
		errors: []error{nil},
	},
	{
		goodPattern: false,
		comments: []string{
			"go.two-fer.some_comment",
			"go.two-fer.some_comment_2",
			"go.two-fer.some_comment_3",
		},
		result: Result{
			Status: DisapproveWithComment,
			Comments: []string{
				"go.two-fer.some_comment",
				"go.two-fer.some_comment_2",
				"go.two-fer.some_comment_3",
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
		suggs.SetSeverity(test.severity)
		for _, comment := range test.comments {
			suggs.AppendUnique(comment)
		}
		for _, err := range test.errors {
			suggs.ReportError(err)
		}

		res := getResult(test.goodPattern, suggs)
		assert.Equal(t, test.result, res)
	}
}

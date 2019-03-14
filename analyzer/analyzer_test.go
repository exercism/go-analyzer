package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var getResultTests = []struct {
	goodPattern bool
	comments    []string
	result      Result
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
}

func Test_getResult(t *testing.T) {
	for _, test := range getResultTests {
		res := getResult(test.goodPattern, test.comments)
		assert.Equal(t, test.result, res)
	}
}

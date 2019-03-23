package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"testing"

	"github.com/exercism/go-analyzer/analyzer"
	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
)

// Tests contains the test cases.
var Tests http.FileSystem = http.Dir("tests")

// var runOnly = ""

func TestAnalyze(t *testing.T) {
	exercises, err := ExercisesWithTests()
	if err != nil {
		t.Fatal(err)
	}

	for _, exercise := range exercises {
		paths, err := ExerciseTests(exercise)
		if err != nil {
			t.Error(err)
			continue
		}

		for _, dir := range paths {
			// if runOnly != "" && runOnly != dir {
			// 	continue
			// }
			res := analyzer.Analyze(exercise, dir)
			expected, err := GetExpected(dir)
			if err != nil {
				t.Errorf("error getting TestResult for path %s: %s", dir, err)
			}

			assert.Equal(t, expected.Status, res.Status,
				fmt.Sprintf("Wrong status on %s (severity: %d, rating: %.2f)", dir, res.Severity, res.Rating))

			checkContains(t, expected.Comments, res.Comments, "Missing comment", dir)
			checkContains(t, res.Comments, expected.Comments, "Additional comment", dir)

			for _, err := range res.Errors {
				assert.Contains(t, expected.Errors, err, "unexpected error analyzing the solution %s: %s", dir, err)
			}
			for _, expError := range expected.Errors {
				assert.Contains(t, res.Errors, expError, "missing error analyzing the solution %s: %s", dir, err)
			}
		}
	}
}

func checkContains(t *testing.T, search, container []sugg.Comment, message, dir string) {
	for _, comment := range search {
		var (
			diff     string
			err      error
			contains = sugg.Contains(container, comment)
			msg      = message
		)
		if !contains {
			cmt := getCommentIdOnly(container, comment.ID())
			if cmt != nil {
				msg = "Different parameters on comment"
				diff, err = commentDiff(comment, cmt)
				if err != nil {
					t.Error(err)
				}
			}
		}
		assert.True(t, contains, fmt.Sprintf("%s `%s` on %s\n%s", msg, comment.ID(), dir, diff))
	}
}

// ExercisesWithTests returns a list of exercise slugs for which tests are provided.
func ExercisesWithTests() ([]string, error) {
	return analyzer.GetDirs(".", Tests)
}

// ExerciseTests returns a list of paths containing tests for given exercise.
func ExerciseTests(exercise string) ([]string, error) {
	paths, err := analyzer.GetDirs(exercise, Tests)
	if err != nil {
		return nil, err
	}

	for i, dir := range paths {
		paths[i] = path.Join("tests", exercise, dir)
	}
	return paths, err
}

// GetExpected returns the content of the `test.json` file in given path.
func GetExpected(dir string) (*analyzer.Result, error) {
	bytes, err := ioutil.ReadFile(path.Join(dir, "expected.json"))
	if err != nil {
		return nil, err
	}

	// transforming to struct and back to json to eliminate different formatting
	var res = unmarshalResult{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		return nil, err
	}

	// bytes, err = toJson(res)
	result := &analyzer.Result{
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
	Status   analyzer.Status `json:"status"`
	Comments []interface{}   `json:"comments"`
	Errors   []string        `json:"errors,omitempty"`
	Severity int             `json:"-"`
}

func commentDiff(expected, got sugg.Comment) (string, error) {
	expectedB, err := json.MarshalIndent(expected, "", "\t")
	if err != nil {
		return "", err
	}
	gotB, err := json.MarshalIndent(got, "", "\t")
	if err != nil {
		return "", err
	}
	return getDiff(expectedB, gotB), nil
}

func getDiff(expected, got []byte) string {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(expected)),
		B:        difflib.SplitLines(string(got)),
		FromFile: "Expected",
		ToFile:   "Got",
		Context:  0,
	}
	text, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		return fmt.Sprintf("error while diffing strings: %s", err)
	}
	return text
}

func getCommentIdOnly(comments []sugg.Comment, id string) sugg.Comment {
	for _, cmt := range comments {
		if cmt.ID() == id {
			return cmt
		}
	}
	return nil
}

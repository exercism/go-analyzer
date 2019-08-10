package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"testing"

	"github.com/exercism/go-analyzer/analyzer"
	"github.com/exercism/go-analyzer/assets"
	"github.com/exercism/go-analyzer/suggester/sugg"
	"github.com/logrusorgru/aurora"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
	"github.com/tehsphinx/astrav"
)

// Tests contains the test cases.
var Tests http.FileSystem = http.Dir("tests")

var runOnly = ""

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
			if runOnly != "" && runOnly != dir {
				continue
			}

			t.Run(dir[6:], func(t *testing.T) {
				res := analyzer.Analyze(exercise, dir)

				// if a specific exercise is set print ast for orientation while implementing
				if runOnly != "" {
					if err := printAST(dir); err != nil {
						t.Fatal(err)
					}
				}

				expected, err := analyzer.GetResultFromFile(dir)
				if err != nil {
					t.Errorf("error getting TestResult for path %s: %s", dir, err)
				}

				var fail bool
				if !assert.Equal(t, expected.Status, res.Status,
					fmt.Sprintf("Wrong status on %s/solution.go:1\n\t-> severity: %d, rating: %.2f", dir, res.Severity, res.Rating)) {
					fail = true
				}

				if checkContains(t, expected.Comments, res.Comments, "Missing comment", dir) {
					fail = true
				}
				if checkContains(t, res.Comments, expected.Comments, "Additional comment", dir) {
					fail = true
				}

				if checkContainsError(t, expected.Errors, res.Errors, dir) {
					fail = true
				}

				if fail {
					diff, err := resultDiff(*expected, res)
					if err != nil {
						t.Error(err)
					}
					t.Errorf("Diff on %s\n%s", dir, diff)
				}
			})
		}
	}
}

func checkContains(t *testing.T, search, container []sugg.Comment, message, dir string) (fail bool) {
	for _, comment := range search {
		var (
			diff     string
			err      error
			contains = sugg.Contains(container, comment)
			msg      = message
		)
		if !contains {
			cmt := getCommentIDOnly(container, comment.ID())
			if cmt != nil {
				msg = "Different parameters on comment"
				diff, err = commentDiff(comment, cmt)
				if err != nil {
					fail = true
					t.Error(err)
				}
			}
		}
		if !assert.True(t, contains, fmt.Sprintf("%s `%s` \n\t-> on %s/solution.go:1\n%s", msg, comment.ID(), dir, diff)) {
			fail = true
		}
	}
	return fail
}

func checkContainsError(t *testing.T, expected, got []string, dir string) (fail bool) {
	for _, err := range expected {
		var found bool
		for _, gotErr := range got {
			if strings.Contains(gotErr, err) {
				found = true
			}
		}
		if !assert.True(t, found, "missing error analyzing the solution %s: %s", dir, err) {
			fail = true
		}
	}

	for _, gotErr := range got {
		var found bool
		for _, err := range expected {
			if strings.Contains(gotErr, err) {
				found = true
			}
		}
		if !assert.True(t, found, "unexpected error analyzing the solution %s: %s", dir, gotErr) {
			fail = true
		}
	}
	return fail
}

// ExercisesWithTests returns a list of exercise slugs for which tests are provided.
func ExercisesWithTests() ([]string, error) {
	return assets.GetDirs(".", Tests)
}

// ExerciseTests returns a list of paths containing tests for given exercise.
func ExerciseTests(exercise string) ([]string, error) {
	paths, err := assets.GetDirs(exercise, Tests)
	if err != nil {
		return nil, err
	}

	for i, dir := range paths {
		paths[i] = path.Join("tests", exercise, dir)
	}
	return paths, err
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

func resultDiff(expected, got analyzer.Result) (string, error) {
	expB, err := json.MarshalIndent(expected, "", "\t")
	if err != nil {
		return "", err
	}
	gotB, err := json.MarshalIndent(got, "", "\t")
	if err != nil {
		return "", err
	}
	diff := getDiff(expB, gotB)
	return diff, nil
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

func getCommentIDOnly(comments []sugg.Comment, id string) sugg.Comment {
	for _, cmt := range comments {
		if cmt.ID() == id {
			return cmt
		}
	}
	return nil
}

func printAST(dir string) error {
	solution, err := analyzer.LoadPackage(dir)
	if err != nil {
		return err
	}
	solution.Walk(func(node astrav.Node) bool {
		src := node.GetSourceString()
		if i := strings.Index(src, "\n"); i != -1 {
			src = src[:i]
		}
		fmt.Printf("%s%s\t\t%s\n", strings.Repeat("\t", node.Level()), aurora.Cyan(fmt.Sprintf("%T", node)), src)
		return true
	})
	return nil
}

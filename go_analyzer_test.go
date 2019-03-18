package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"testing"

	"github.com/exercism/go-analyzer/analyzer"
	"github.com/stretchr/testify/assert"
)

// Tests contains the test cases.
var Tests http.FileSystem = http.Dir("tests")

// TestCase defines the structure for a test case.
// A test case is a folder containing a solution and a file with the `test.json`
// containing the TestCase structure.
type TestCase struct {
	ExpectedStatus      analyzer.Status `json:"expected_status"`
	ExpectedComments    []string        `json:"expected_comments"`
	NotExpectedComments []string        `json:"not_expected_comments"`
}

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
			for _, err := range res.Errors {
				t.Errorf("error analyzing the solution %s: %s", dir, err)
			}

			bytes, err := toJson(res)
			if err != nil {
				t.Errorf("error transforming to json for path %s: %s", dir, err)
			}

			expected, err := GetExpected(dir)
			if err != nil {
				t.Errorf("error getting TestResult for path %s: %s", dir, err)
			}
			assert.Equal(t, string(bytes), expected, "result is not as expected on %s", dir)
		}
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
func GetExpected(dir string) (string, error) {
	bytes, err := ioutil.ReadFile(path.Join(dir, "expected.json"))
	if err != nil {
		return "", err
	}

	// transforming to struct and back to json to eliminate different formatting
	var res = analyzer.Result{}
	if err := json.Unmarshal(bytes, &res); err != nil {
		return string(bytes), err
	}

	bytes, err = toJson(res)
	return string(bytes), err
}

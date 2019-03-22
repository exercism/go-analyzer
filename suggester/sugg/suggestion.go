package sugg

import (
	"encoding/json"
	"fmt"
)

// Comment defines a suggestion. A comment to the student.
type Comment interface {
	compareString() string
	setSeverity(int)

	// ID returns the comment identifier. e.g. `go.two_fer.missing_share_with_function`
	ID() string

	// Severity reports the severity of the comment.
	Severity() int
}

// Contains reports if the list of comments includes a certain comment.
func Contains(comments []Comment, comment Comment) bool {
	for _, cmt := range comments {
		if cmt.compareString() == comment.compareString() {
			return true
		}
	}
	return false
}

// NewComment creates a new comment
func NewComment(s string) Comment {
	return &comment{
		comment: s,
	}
}

type comment struct {
	comment  string
	severity int
}

// ID returns the comment identifier.
func (s *comment) ID() string {
	return s.comment
}

func (s *comment) compareString() string {
	return s.ID()
}

func (s *comment) setSeverity(severity int) {
	s.severity = severity
}

// Severity reports the severity of the comment.
func (s *comment) Severity() int {
	return s.severity
}

// MarshalJSON converts the comment to json.
func (s *comment) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.comment)
}

// UnmarshalJSON converts json to comment.
func (s *comment) UnmarshalJSON(data []byte) error {
	var v string
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*s = comment{
		comment: v,
	}
	return nil
}

// NewPlaceholderComment creates a new comment with placeholder(s).
func NewPlaceholderComment(comment string, params map[string]string) Comment {
	return &placeholderComment{
		comment: comment,
		params:  params,
	}
}

type placeholderComment struct {
	comment  string
	params   map[string]string
	severity int
}

// ID returns the comment identifier.
func (s *placeholderComment) ID() string {
	return s.comment
}

func (s *placeholderComment) compareString() string {
	return fmt.Sprintf("%s;%v", s.comment, s.params)
}

func (s *placeholderComment) setSeverity(severity int) {
	s.severity = severity
}

// Severity reports the severity of the comment.
func (s *placeholderComment) Severity() int {
	return s.severity
}

// MarshalJSON converts the placeholderComment to json.
func (s *placeholderComment) MarshalJSON() ([]byte, error) {

	return json.Marshal(struct {
		Comment string            `json:"comment"`
		Params  map[string]string `json:"params"`
	}{
		Comment: s.comment,
		Params:  s.params,
	})
}

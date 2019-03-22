package sugg

import (
	"encoding/json"
	"fmt"
)

type suggestion struct {
	comment  Comment
	severity int
}

// Comment defines a suggestion. A comment to the student.
type Comment interface {
	compareString() string

	// ID returns the comment identifier. e.g. `go.two_fer.missing_share_with_function`
	ID() string
}

func newComment(s string) *comment {
	return (*comment)(&s)
}

type comment string

// ID returns the comment identifier.
func (s *comment) ID() string {
	return string(*s)
}

func (s *comment) compareString() string {
	return s.ID()
}

// MarshalJSON converts the comment to json.
func (s *comment) MarshalJSON() ([]byte, error) {
	return json.Marshal(*s)
}

// UnmarshalJSON converts json to comment.
func (s *comment) UnmarshalJSON(data []byte) error {
	var v string
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	*s = comment(v)
	return nil
}

type placeholderComment struct {
	Comment string            `json:"comment"`
	Params  map[string]string `json:"params"`
}

// ID returns the comment identifier.
func (s *placeholderComment) ID() string {
	return s.Comment
}

func (s *placeholderComment) compareString() string {
	return fmt.Sprintf("%s;%v", s.Comment, s.Params)
}

// MarshalJSON converts the placeholderComment to json.
func (s *placeholderComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(*s)
}

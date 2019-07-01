package sugg

import (
	"encoding/json"
	"fmt"
	"log"
)

// Category defines an enumeration of comment categories
type Category string

// Category contants
const (
	CtgTodo        Category = "todo"
	CtgImprovement Category = "improvement"
	CtgThought     Category = "thought"
	CtgBlock       Category = "block"
)

// Comment defines a suggestion. A comment to the student.
type Comment interface {
	compareString() string
	setSeverity(int)

	// ID returns the comment identifier. e.g. `go.two-fer.missing_share_with_function`
	ID() string

	// Severity reports the severity of the comment.
	Severity() int

	// ToString returns the final comment the way it will be provided to the student.
	// If not already cached this involves pulling the comment from the git repository where it is located.
	ToString() string

	// Category returns the category of the comment.
	Category() Category
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
		baseCmt: baseCmt{
			comment: s,
		},
	}
}

// NewBlockComment creates a new block comment
func NewBlockComment(s string) Comment {
	return &comment{
		baseCmt: baseCmt{
			comment:  s,
			category: CtgBlock,
		},
	}
}

type comment struct {
	baseCmt
}

// ToString returns the final comment text
func (s *comment) ToString() string {
	return s.toString(s)
}

// MarshalJSON converts the comment to json.
func (s *comment) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.comment)
}

// NewPlaceholderComment creates a new comment with placeholder(s).
func NewPlaceholderComment(comment string, params map[string]string) Comment {
	return &placeholderComment{
		baseCmt: baseCmt{
			comment: comment,
		},
		params: params,
	}
}

type placeholderComment struct {
	baseCmt

	params map[string]string
}

func (s *placeholderComment) compareString() string {
	return fmt.Sprintf("%s;%v", s.comment, s.params)
}

// ToString returns the final comment text
func (s *placeholderComment) ToString() string {
	return s.toString(s)
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

// baseComment
type baseCmt struct {
	comment  string
	severity int
	category Category
}

// ID returns the comment identifier.
func (s *baseCmt) ID() string {
	return s.comment
}

func (s *baseCmt) compareString() string {
	return s.ID()
}

func (s *baseCmt) setSeverity(severity int) {
	s.severity = severity
	if s.category == "" {
		s.category = sevToCat(severity)
	}
}

// Severity reports the severity of the comment.
func (s *baseCmt) Severity() int {
	return s.severity
}

// Category reports the category of the comment.
func (s *baseCmt) Category() Category {
	return s.category
}

func (*baseCmt) toString(s Comment) string {
	txt, err := comments.GetCommentText(s)
	if err != nil {
		log.Println(err)
	}
	return txt
}

func sevToCat(severity int) Category {
	switch {
	case 5 <= severity:
		return CtgTodo
	case severity == 0:
		return CtgThought
	}
	return CtgImprovement
}

package sugg

// general comments
const (
	SyntaxError      = "go.general.syntax_error"
	CommentSection   = "go.general.section_about_comments"
	LenOfStringEqual = "go.general.taking_lenght_of_string_to_check_empty"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	SyntaxError:      5,
	CommentSection:   0,
	LenOfStringEqual: 2,
}

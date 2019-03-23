package sugg

// general comments
const (
	SyntaxError       = "go.general.syntax_error"
	CommentSection    = "go.general.section_about_comments"
	LenOfStringEqual  = "go.general.taking_length_of_string_to_check_empty"
	MainFunction      = "go.general.remove_main_function_and_correct_package_name"
	GoFmt             = "go.general.gofmt_not_used"
	GoLint            = "go.general.golint_not_satisfied"
	LengthSmallerZero = "go.general.length_smaller_zero_impossible"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	SyntaxError:       5,
	CommentSection:    0,
	LenOfStringEqual:  2,
	MainFunction:      5,
	GoFmt:             5,
	GoLint:            5,
	LengthSmallerZero: 2,
}

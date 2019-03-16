package twofer

import "github.com/exercism/go-analyzer/suggester/sugg"

// exercise comments
const (
	CommentSection         = sugg.CommentSection
	MissingShareWith       = "go.two_fer.missing_share_with_function"
	StringsJoin            = "go.two_fer.strings_join_used_for_concatenation"
	PlusUsed               = "go.two_fer.plus_used_for_concatenation"
	MinimalConditional     = "go.two_fer.find_minimal_conditional"
	UseStringPH            = "go.two_fer.use_fmt_placeholder_for_string"
	StubComments           = "go.two_fer.replace_stub_comments"
	MissingPackageComment  = "go.two_fer.missing_package_comment"
	MissingFunctionComment = "go.two_fer.missing_function_comment"
	WrongPackageComment    = "go.two_fer.wrong_package_comment"
	WrongFunctionComment   = "go.two_fer.wrong_function_comment"
	GeneralizeName         = "go.two_fer.work_with_any_provided_name"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	CommentSection:         0,
	MissingShareWith:       5,
	StringsJoin:            5,
	PlusUsed:               0,
	MinimalConditional:     5,
	UseStringPH:            2,
	StubComments:           5,
	MissingPackageComment:  1,
	MissingFunctionComment: 2,
	WrongPackageComment:    2,
	WrongFunctionComment:   2,
	GeneralizeName:         5,
}

package twofer

import "github.com/exercism/go-analyzer/suggester/sugg"

// exercise comments
const (
	MissingEntryFunc       = "go.two-fer.missing_share_with_function"
	FuncSignatureChanged   = "go.two-fer.sharewith_signature_changed"
	CommentSection         = sugg.CommentSection
	StringsJoin            = "go.two-fer.strings_join_used_for_concatenation"
	StringsBuilder         = "go.two-fer.strings_builder_used_for_concatenation"
	PlusUsed               = "go.two-fer.plus_used_for_concatenation"
	MinimalConditional     = "go.two-fer.find_minimal_conditional"
	UseStringPH            = "go.two-fer.use_fmt_placeholder_for_string"
	StubComments           = "go.two-fer.replace_stub_comments"
	MissingPackageComment  = "go.two-fer.missing_package_comment"
	MissingFunctionComment = "go.two-fer.missing_function_comment"
	WrongPackageComment    = "go.two-fer.wrong_package_comment"
	WrongFunctionComment   = "go.two-fer.wrong_function_comment"
	GeneralizeName         = "go.two-fer.work_with_any_provided_name"
	ExtraNameVar           = "go.two-fer.extra_name_variable_created"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	CommentSection:         0,
	MissingEntryFunc:       5,
	StringsJoin:            5,
	StringsBuilder:         5,
	PlusUsed:               0,
	MinimalConditional:     5,
	UseStringPH:            2,
	StubComments:           5,
	MissingPackageComment:  1,
	MissingFunctionComment: 2,
	WrongPackageComment:    2,
	WrongFunctionComment:   2,
	GeneralizeName:         5,
	FuncSignatureChanged:   5,
	ExtraNameVar:           1,
}

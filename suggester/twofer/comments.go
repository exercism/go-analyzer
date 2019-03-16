package twofer

// exercise comments
const (
	MissingShareWith   = "go.two_fer.missing_share_with_function"
	StringsJoin        = "go.two_fer.strings_join_used_for_concatenation"
	PlusUsed           = "go.two_fer.plus_used_for_concatenation"
	MinimalConditional = "go.two_fer.find_minimal_conditional"
	UseStringPH        = "go.two_fer.use_fmt_placeholder_for_string"
	StubComments       = "go.two_fer.replace_stub_comments"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	StringsJoin:        5,
	PlusUsed:           0,
	MinimalConditional: 5,
	UseStringPH:        2,
	StubComments:       5,
}

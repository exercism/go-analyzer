package sugg

// general comments
const (
	AvoidInit          = "go.general.avoid_init_function"
	AvoidPrint         = "go.general.avoid_printing_and_logging"
	CustomErrorCreated = "go.general.custom_error_created"
	ErrorfWithoutParam = "go.general.fmt_errorf_without_parameter"
	GoFmt              = "go.general.gofmt_not_used"
	GoLint             = "go.general.golint_not_satisfied"
	LengthSmallerZero  = "go.general.length_smaller_zero_impossible"
	OmittedErrorMsg    = "go.general.omitted_error_message"
	ExtraFunction      = "go.general.remove_extra_function"
	ExtraVar           = "go.general.remove_extra_variable"
	MainFunction       = "go.general.remove_main_function_and_correct_package_name"
	CommentSection     = "go.general.section_about_comments"
	StringsCompare     = "go.general.strings_compare_used"
	TrimSpaceUsed      = "go.general.strings_trim_space_used"
	SyntaxError        = "go.general.syntax_error"
	LenOfStringEqual   = "go.general.taking_length_of_string_to_check_empty"
	UseVarAssignment   = "go.general.use_variable_assignment"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	AvoidInit:          5,
	AvoidPrint:         5,
	SyntaxError:        5,
	CommentSection:     0,
	LenOfStringEqual:   2,
	MainFunction:       5,
	GoFmt:              5,
	GoLint:             5,
	LengthSmallerZero:  2,
	ExtraVar:           3,
	UseVarAssignment:   1,
	ExtraFunction:      5,
	OmittedErrorMsg:    3,
	ErrorfWithoutParam: 2,
	CustomErrorCreated: 0,
	TrimSpaceUsed:      0,
	StringsCompare:     3,
}

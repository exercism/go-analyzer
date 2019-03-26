package hamming

// exercise comments
const (
	MissingMainFunc      = "go.hamming.missing_distance_function"
	FuncSignatureChanged = "go.hamming.distance_signature_changed"
	DeclareWhenNeeded    = "go.hamming.declare_variable_when_needed_not_start_of_function"
	ErrorMsgFormat       = "go.hamming.error_msg_not_capitalized_nor_punctuated"
	IncreaseOperator     = "go.hamming.use_increase_operator"
	InvertIf             = "go.hamming.invert_if_for_happy_path_on_left"
	NakedReturns         = "go.hamming.use_of_naked_returns"
	OmittedErrorMsg      = "go.hamming.omitted_error_message"
	ZeroValueOnErr       = "go.hamming.use_zero_value_for_other_return_params_on_error"
	RuneToByte           = "go.hamming.lossy_rune_to_byte_conversion"
	MultipleStringConv   = "go.hamming.multiple_string_conversions"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	MissingMainFunc:      5,
	FuncSignatureChanged: 5,
	DeclareWhenNeeded:    1,
	ErrorMsgFormat:       1,
	IncreaseOperator:     2,
	InvertIf:             3,
	NakedReturns:         3,
	OmittedErrorMsg:      3,
	ZeroValueOnErr:       2,
	RuneToByte:           3,
	MultipleStringConv:   3,
}

package hamming

// exercise comments
const (
	MissingMainFunc      = "go.hamming.missing_distance_function"
	FuncSignatureChanged = "go.hamming.distance_signature_changed"
	MixtureRunesBytes    = "go.hamming.mixture_of_runes_and_bytes"
	RuneByteIndex        = "go.hamming.rune_index_used_with_bytes"
	DeclareWhenNeeded    = "go.hamming.declare_variable_when_needed_not_start_of_function"
	ErrorMsgFormat       = "go.hamming.error_msg_not_capitalized_nor_punctuated"
	IncreaseOperator     = "go.hamming.use_increase_operator"
	InvertIf             = "go.hamming.invert_if_for_happy_path_on_left"
	NakedReturns         = "go.hamming.use_of_naked_returns"
	ZeroValueOnErr       = "go.hamming.use_zero_values_on_error_return"
	RuneToByte           = "go.hamming.lossy_rune_to_byte_conversion"
	ToStringConversion   = "go.hamming.rune_or_byte_to_string_conversion"
	StringsSplitUsed     = "go.hamming.strings_split_used"
	ComparingBytes       = "go.hamming.comparing_bytes"
	DefineEmptyErr       = "go.hamming.defining_empty_error"
	ReturnOnError        = "go.hamming.return_on_error_case"
	CaseInsensitive      = "go.hamming.hamming_is_case_sensitive"
	TrimSpaceUsed        = "go.hamming.strings_trim_space_used"
	ExtraIfStringsEmpty  = "go.hamming.extra_if_strings_empty"
	MinSliceRuneConv     = "go.hamming.minimize_slice_of_rune_conversions"
	CompBytesInDisguise  = "go.hamming.comparing_bytes_disguised_as_runes"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	MissingMainFunc:      5,
	FuncSignatureChanged: 5,
	MixtureRunesBytes:    3,
	RuneByteIndex:        2,
	DeclareWhenNeeded:    1,
	ErrorMsgFormat:       1,
	IncreaseOperator:     2,
	InvertIf:             3,
	NakedReturns:         3,
	ZeroValueOnErr:       2,
	RuneToByte:           3,
	ToStringConversion:   3,
	StringsSplitUsed:     5,
	ComparingBytes:       0,
	DefineEmptyErr:       1,
	ReturnOnError:        3,
	CaseInsensitive:      3,
	TrimSpaceUsed:        0,
	ExtraIfStringsEmpty:  0,
	MinSliceRuneConv:     2,
	CompBytesInDisguise:  5,
}

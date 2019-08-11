package scrabble

// exercise comments
const (
	MissingEntryFunc     = "go.scrabble.missing_share_with_function"
	FuncSignatureChanged = "go.scrabble.sharewith_signature_changed"
	Regex                = "go.scrabble.static_regex_in_func"
	GoRoutines           = "go.scrabble.using_goroutines"
	IfsToSwitch          = "go.scrabble.transform_ifs_to_switch"
	MoveMap              = "go.scrabble.static_map_in_func"
	SliceRuneConv        = "go.scrabble.no_slice_rune_conversion_before_loop"
	UnicodeLoop          = "go.scrabble.use_unicode_instead_strings_in_loop"
	Unicode              = "go.scrabble.use_unicode_in_loop_instead_strings_before_loop"
	MultipleLoops        = "go.scrabble.multiple_loops"
	MapRune              = "go.scrabble.use_map_rune"
	LoopRuneNotByte      = "go.scrabble.iterate_runes_not_bytes"
	TypeConversion       = "go.scrabble.unnecessary_type_conversion"
	RegexChallenge       = "go.scrabble.challenge"
	TrySwitch            = "go.scrabble.try_switch"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	MissingEntryFunc:     5,
	FuncSignatureChanged: 5,
	Regex:                5,
	GoRoutines:           5,
	IfsToSwitch:          5,
	MoveMap:              5,
	SliceRuneConv:        2,
	UnicodeLoop:          1,
	Unicode:              1,
	MultipleLoops:        3,
	MapRune:              3,
	LoopRuneNotByte:      3,
	TypeConversion:       2,
	RegexChallenge:       0,
	TrySwitch:            0,
}

package raindrops

import "github.com/exercism/go-analyzer/suggester/sugg"

// exercise comments
const (
	MissingEntryFunc     = "go.raindrops.missing_share_with_function"
	FuncSignatureChanged = "go.raindrops.sharewith_signature_changed"
	AllCombinations      = "go.raindrops.all_combinations_implemented"
	ConcatFMT            = "go.raindrops.concatenation_via_fmt"
	RemoveExtraBool      = "go.raindrops.remove_extra_bool"
	LoopMap              = "go.raindrops.loop_over_a_map"
	ExtensiveFor         = "go.raindrops.extensive_for_loop"
	UseItoa              = "go.raindrops.use_strconv_itoa"
	PlusEqual            = "go.raindrops.plus_equal_comment"
	MultipleLoops        = "go.raindrops.multiple_loops"
	StringBuilder        = "go.raindrops.string_builder_used"
)

// Severity defines how severe a comment is. A sum over all comments of 5 means no approval.
// The maximum for a single comment is 5. A comment with that severity will block approval.
// When assigning the severity a good guideline is to ask: How many comments of similar severity
// should block approval?
// We can be very strict on automated comments since the student has a very fast feedback loop.
var severity = map[string]int{
	MissingEntryFunc:     5,
	FuncSignatureChanged: 5,
	AllCombinations:      5,
	ConcatFMT:            2,
	RemoveExtraBool:      2,
	LoopMap:              5,
	ExtensiveFor:         3,
	UseItoa:              2,
	PlusEqual:            0,
	MultipleLoops:        5,
	StringBuilder:        2,
	sugg.ExtraFunction:   1,
}

package sugg

import (
	"bytes"
	"fmt"
	"go/format"

	"github.com/pmezard/go-difflib/difflib"
)

func fmtCode(files map[string][]byte) string {
	for _, file := range files {
		file = bytes.Replace(file, []byte{'\r', '\n'}, []byte{'\n'}, -1)
		f, err := format.Source(file)
		if err != nil {
			return fmt.Sprintf("code fails to format with error: %s\n", err)
		}
		if string(f) != string(file) && string(f) != string(append(file, '\n')) {
			return getDiff(file, f)
		}
	}
	return ""
}

func getDiff(current, formatted []byte) string {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(string(current)),
		B:        difflib.SplitLines(string(formatted)),
		FromFile: "Current",
		ToFile:   "Formatted",
		Context:  0,
	}
	text, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		return fmt.Sprintf("error while diffing strings: %s", err)
	}
	return text
}

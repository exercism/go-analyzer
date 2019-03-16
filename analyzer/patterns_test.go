package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatternDirs(t *testing.T) {
	dirs, err := PatternDirs("two-fer")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, dirs)
}

func TestGetDirs(t *testing.T) {
	dirs, err := GetDirs("two-fer", Patterns)
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, dirs)
}

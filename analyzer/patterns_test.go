package analyzer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPatterns(t *testing.T) {
	dirs, err := PatternDirs("two-fer")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, dirs)
}

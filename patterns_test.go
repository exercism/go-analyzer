package main

import (
	"testing"

	"github.com/exercism/go-analyzer/assets"
	"github.com/stretchr/testify/assert"
)

func TestPatternDirs(t *testing.T) {
	patterns, err := assets.LoadPatterns("two-fer")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, patterns)
}

func TestGetDirs(t *testing.T) {
	dirs, err := assets.GetDirs("two-fer", assets.Patterns)
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, dirs)
}

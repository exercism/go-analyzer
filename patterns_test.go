package main

import (
	"testing"

	"github.com/exercism/go-analyzer/analyzer"
	"github.com/stretchr/testify/assert"
)

func TestPatternDirs(t *testing.T) {
	dirs, err := analyzer.PatternDirs("two-fer")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, dirs)
}

func TestGetDirs(t *testing.T) {
	dirs, err := analyzer.GetDirs("two-fer", analyzer.Folder)
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, dirs)
}

# Contributing

Thanks so much for wanting to add to the Exercism automated analyzer for the
Go track.

## Table of Contents

<!-- vim-markdown-toc GFM -->

- [Adding a New Exercise to the Analyzer](#adding-a-new-exercise-to-the-analyzer)
  - [Overview](#overview)

<!-- vim-markdown-toc -->

## Adding a New Exercise to the Analyzer

### Overview

This document describes the steps needed to analyze a Go exercise with this
analyzer.

The summary of the steps are as follows:

1. Add the canonical solution for the exercise in a numbered directory under
   `/tests/<exercise_name>`
2. Add the expected analyzer output for the canonical solution to `expected.json`
3. Copy the canonical solution as a pattern for the exercise in a numbered
   directory under `/assets/patterns/<exercise_name>`
4. Run `go test go_analyzer_test.go` to ensure that output is as expected

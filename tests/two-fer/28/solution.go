package main

import "fmt"

// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
//package twofer

// This function grabs one integer and one string, then passes it through a loop.
func ShareWith(x int, name string) {
	for x == 1 {
		x = x + 1
		if x == 3 {
			break
		} else if name != "" {
			fmt.Println("One for", name, ", one for me.")
		} else {
			fmt.Println("One for you, one for me")
		}
	}
}
func main() {
	ShareWith(1, "alice")
	ShareWith(1, "")
}

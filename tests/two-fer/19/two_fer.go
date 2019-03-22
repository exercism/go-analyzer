// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package twofer should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package main

//package twofer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text:")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	yourName := ShareWith(text)

	fmt.Printf("One for %s, one for me. \n", yourName)

}

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {

	if name == "" {
		return "you"
	} else {
		return name
	}

}

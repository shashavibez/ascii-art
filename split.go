package main

import "strings"

// Split splits the input string by the literal \n sequence
// and returns a slice of strings.
func Split(input string) []string {
	return strings.Split(input, `\n`)
}

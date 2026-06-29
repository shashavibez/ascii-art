package main

import "strings"

// RenderLine takes a single line of text and renders it
// as ASCII art using the provided charMap.
func RenderLine(input string, banner map[rune][]string) []string {
	style := make([]string, 8)
	for row := range 8 {
		var styleBulider strings.Builder
		for _, i := range input {
			styleBulider.WriteString(banner[i][row])
		}
		style[row] = styleBulider.String()
	}
	return style
}

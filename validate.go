package main

import "errors"

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < rune(32) || char > 126 {
			return char, errors.New("unsupported input")
		}
	}
	return 0, nil
}

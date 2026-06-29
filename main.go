package main

import (
	"fmt"
	"os"
)

func main() {
	err := ValidateArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	input := os.Args[1]

	bannerFile := "standard.txt"
	if len(os.Args) == 3 {
		bannerFile = os.Args[2] + ".txt"
	}

	charMap, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
		os.Exit(1)
	}

	err = ValidateChars(input, charMap)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if input == "" {
		fmt.Println()
		os.Exit(0)
	}

	
	output := Generate(input, charMap)
	fmt.Print(output)
}

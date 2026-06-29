package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// validate arguments
	if len(os.Args) < 2 || len(os.Args) > 3 {
		log.Fatal("usage: go run <string> <Banner style(optional)>")
		return
	}

	input := os.Args[1]

	// default banner is standard
	bannerFile := "standard.txt"
	if len(os.Args) == 3 {
		// if banner argument does not end with ".txt", add it
		if !(strings.HasSuffix(os.Args[2], ".txt")) {
			bannerFile = os.Args[2] + ".txt"
		} else {
			// use banner argument if it ends with ".txt"
			bannerFile = os.Args[2]
		}

	}

	// load the banner file
	charMap, err := LoadBanner(bannerFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading banner: %v\n", err)
		os.Exit(1)
	}

	// handle empty input
	if input == "" {
		fmt.Println()
		os.Exit(0)
	}

	// generate and print the ASCII art
	output := GenerateArt(input, charMap)
	fmt.Print(output)
}

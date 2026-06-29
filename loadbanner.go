package main

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
)

func LoadBanner(file string) (map[rune][]string, error) {
	var result []string

	var bannerFile string
	// checks if a file path is absolute or not. E.g. c:\... or /...
	//  helps support file created for Testing purpose.
	if filepath.IsAbs(file) {
		// mostly for test files
		bannerFile = file
	} else {
		// for our core application logic
		bannerFile = "banner_files/" + file
	}

	filename, err := os.Open(bannerFile) // open file from banner folder
	if err != nil {
		return nil, errors.New("file does not exist or error opening file")
	}
	defer filename.Close()

	scanner := bufio.NewScanner(filename)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Error scanning file")
	}

	if len(result) == 0 {
		return nil, errors.New("empty file")
	}

	if len(result) != 855 {
		return nil, errors.New("length of banner file must not be less than 855")
	}

	if len(result)%9 != 0 {
		return nil, errors.New("incomplete file")
	}

	font := make(map[rune][]string)

	for i := 1; i+7 < len(result); i += 9 {
		block := result[i : i+8]
		char := rune((i / 9) + 32)

		if len(block) != 8 {
			return nil, errors.New("invalid banner")
		}

		font[char] = block
	}

	if len(font) == 0 {
		return nil, errors.New("invalid banner")
	}

	return font, nil
}

package main

import (
	"os"
	"strings"
)

const (
	charsPerBlock = 9   
	charHeight    = 8  
)

func LoadBanner(filename string) (map[rune][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	lines := strings.Split(content, "\n")
	fmt.Println(lines)

	charMap := make(map[rune][]string)

	for i := 32; i <= 126; i++ {
		index := (i-32)*charsPerBlock + 1
		if index+charHeight > len(lines) {
			break
		}
		charMap[rune(i)] = lines[index : index+charHeight]
	}

	return charMap, nil
}

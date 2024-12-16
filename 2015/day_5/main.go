package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", path))
	}
	return strings.Split(string(content), "\n")
}

func atLeastNVowels(str string, n int) bool {
	count := 0
	vowels := "aeiou"
	for _, s := range str {
		if strings.Contains(vowels, string(s)) {
			count++
		}
	}

	if count >= n {
		return true
	}
	return false
}

func main() {
	niceStrings := 0
	lines := ReadFile("test.txt")

	for _, str := range lines {
		fmt.Println(str)
		niceStrings++
	}
}

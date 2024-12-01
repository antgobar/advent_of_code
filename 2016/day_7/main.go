package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func ReadTextFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	return strings.Split(string(content), "\n")
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	reverse_container := make([]rune, len(runes))

	for i, letter := range runes {
		reverse_container[len(s)-1-i] = rune(letter)
	}

	return s == string(reverse_container)
}

func containsValidPalindrome(str string, size int) bool {
	if size <= 0 || len(str) < size {
		return false
	}

	for i := 0; i <= len(str)-size; i++ {
		n_substring := str[i : i+size]
		if isPalindrome(n_substring) && !allCharsTheSame(n_substring) {
			return true
		}
	}
	return false
}

type BracketGroups struct {
	inside  []string
	outside []string
}

func HasPalindrome(strings []string) bool {
	for _, str := range strings {
		if containsValidPalindrome(str, 4) {
			return true
		}
	}
	return false
}

func extractStrings(inputString string) ([]string, []string) {
	reInside := regexp.MustCompile(`\[(.*?)\]`)
	insideBrackets := reInside.FindAllStringSubmatch(inputString, -1)
	var inside []string
	for _, match := range insideBrackets {
		inside = append(inside, match[1])
	}

	reOutside := regexp.MustCompile(`\[[^\[\]]*\]`)
	outsideParts := reOutside.Split(inputString, -1)
	var outside []string
	for _, part := range outsideParts {
		if strings.TrimSpace(part) != "" {
			outside = append(outside, part)
		}
	}
	return inside, outside
}

func allCharsTheSame(strs string) bool {
	first := rune(strs[0])

	for _, s := range strs {
		if first != s {
			return false
		}
	}
	return true
}

func main() {
	data := ReadTextFile("day_7.txt")
	valid := 0

	for _, line := range data {
		inside, outside := extractStrings(line)
		if !HasPalindrome(inside) && HasPalindrome(outside) {
			valid++
		}
	}

	fmt.Println(valid)
}

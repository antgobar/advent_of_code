package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	runes := []rune(s)
	reverse_container := make([]rune, len(runes))

	for i, letter := range runes {
		reverse_container[len(s)-1-i] = rune(letter)
	}

	return s == string(reverse_container)
}

func main() {
	text1 := "abba"
	text2 := "abcd"

	fmt.Println(isPalindrome(text1), isPalindrome(text2))

}

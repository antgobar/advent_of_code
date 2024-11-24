package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadTextFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	return strings.Split(string(content), "\n")
}

func main() {
	data := ReadTextFile("day_6.txt")

	groups := make([]map[string]int, len(data[0]))

	for _, line := range data {
		chars := strings.Split(line, "")
		for i, char := range chars {
			group := groups[i]
			if group == nil {
				group = make(map[string]int)
				groups[i] = group
			}
			group[char]++
		}
	}

	var message1 []string
	var message2 []string

	for _, group := range groups {
		message1 = append(message1, findKeyWithMaxValue(group))
		message2 = append(message2, findKeyWithMinValue(group))
	}

	fmt.Println("Problem 1:", strings.Join(message1, ""))
	fmt.Println("Problem 2:", strings.Join(message2, ""))

}

func findKeyWithMaxValue(m map[string]int) string {
	if len(m) == 0 {
		return ""
	}

	maxVal := 0
	maxKey := ""

	for key, val := range m {
		if val > maxVal {
			maxVal = val
			maxKey = key
		}
	}

	return maxKey
}

func findKeyWithMinValue(m map[string]int) string {
	if len(m) == 0 {
		return ""
	}

	maxVal := 100000000000
	maxKey := ""

	for key, val := range m {
		if val < maxVal {
			maxVal = val
			maxKey = key
		}
	}

	return maxKey
}

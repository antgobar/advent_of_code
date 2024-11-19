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
	return strings.Split(string(content), ", ")
}

func main() {
	ReadTextFile("day_2_test.txt")
}

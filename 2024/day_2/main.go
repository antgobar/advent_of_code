package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file %v", path))
	}

	return strings.Split(string(content), "\n")
}

func ReportIsSafe(report []string) bool {
	return true
}

func main() {
	data := ReadFile("test.txt")

	safeReports := 0

	for _, line := range data {
		fmt.Println(line)
	}

	fmt.Println("Problem 1: ", safeReports)
}

package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFile(path string) string {
	content, _ := os.ReadFile(path)
	data := string(content)
	return strings.Replace(data, "\n", "", -1)
}

func extractMulMatches(line string) []string {
	matcher := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	return matcher.FindAllString(line, -1)
}

func MustConvertStrInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Error parsing string %v as number", num))
	}
	return num
}

func mulCompute(muls []string) int {
	total := 0

	for _, match := range muls {
		matcher := regexp.MustCompile(`\d+`)
		nums := matcher.FindAllString(match, -1)

		total += MustConvertStrInt(nums[0]) * MustConvertStrInt(nums[1])
	}
	return total
}

func cleanDonts(data string) string {
	matcher := regexp.MustCompile(`don't\(([^()]*)\).*?do\(([^()]*)\)`)
	return matcher.ReplaceAllString(data, "")
}

func main() {
	data := ReadFile("input.txt")

	nums1 := extractMulMatches(data)
	totalProblem1 := mulCompute(nums1)
	fmt.Println("Problem 1:", totalProblem1)

	cleaned := cleanDonts(data)
	nums2 := extractMulMatches(cleaned)
	totalProblem2 := mulCompute(nums2)
	fmt.Println("Problem 2:", totalProblem2)
}

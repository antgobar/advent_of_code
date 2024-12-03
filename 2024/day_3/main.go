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

func extractMuls(line string) []string {
	mulMatch := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	return mulMatch.FindAllString(line, -1)
}

func MustStr2Int(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Error parsing string %v as number", num))
	}
	return num
}

func computeMultSum(muls []string) int {
	total := 0

	for _, match := range muls {
		extractedNums := regexp.MustCompile(`\d+`)
		nums := extractedNums.FindAllString(match, -1)

		total += MustStr2Int(nums[0]) * MustStr2Int(nums[1])
	}
	return total
}

func removeDontsToDos(data string) string {
	matcher := regexp.MustCompile(`don't\(([^()]*)\).*?do\(([^()]*)\)`)
	return matcher.ReplaceAllString(data, "")
}

func main() {
	data := ReadFile("input.txt")
	nums1 := extractMuls(data)
	totalProblem1 := computeMultSum(nums1)
	fmt.Println("Problem 1: ", totalProblem1)

	cleaned := removeDontsToDos(data)
	nums2 := extractMuls(cleaned)
	totalProblem2 := computeMultSum(nums2)
	fmt.Println("Problem 2: ", totalProblem2)
}

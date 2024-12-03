package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	content, _ := os.ReadFile(path)
	return strings.Split(string(content), "\n")
}

func extractMuls(line string) []string {
	mulMatch := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	return mulMatch.FindAllString(line, -1)
}

func computeMultSum(muls []string) int {
	total := 0

	for _, match := range muls {
		extractedNums := regexp.MustCompile(`\d+`)
		nums := extractedNums.FindAllString(match, -1)
		firstNum, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(fmt.Sprintf("Error parsing string %v as number", firstNum))
		}
		secondNum, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(fmt.Sprintf("Error parsing string %v as number", secondNum))
		}
		total += firstNum * secondNum
	}
	return total

}

func main() {
	data := ReadFile("input.txt")

	totalProblem1 := 0
	for _, line := range data {
		nums := extractMuls(line)
		totalProblem1 += computeMultSum(nums)
	}
	fmt.Println("Problem 1: ", totalProblem1)
}

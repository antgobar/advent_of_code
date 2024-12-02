package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file %v", path))
	}

	return strings.Split(string(content), "\n")
}

type Report []int

func (report Report) isOrdered() bool {
	if len(report) <= 1 {
		return true
	}

	ascending := true
	descending := true

	for i := 1; i < len(report); i++ {
		if report[i] <= report[i-1] {
			ascending = false
		}
		if report[i] >= report[i-1] {
			descending = false
		}
	}
	return ascending || descending
}

func (report Report) HasMaxDifference(n int) bool {
	for i := 1; i < len(report); i++ {
		numA := float64(report[i])
		numB := float64(report[i-1])

		if int(math.Abs(numA-numB)) > n {
			return false
		}
	}
	return true
}

func ReportIsSafe(line []string) bool {
	var numbers []int

	for _, s := range line {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(fmt.Sprintf("Error converting string %v", num))
		}
		numbers = append(numbers, num)
	}

	report := Report(numbers)

	if !report.isOrdered() {
		return false
	}

	if !report.HasMaxDifference(3) {
		return false
	}
	return true
}

func removeAtIndex(levels []string, index int) []string {
	if index < 0 || index >= len(levels) {
		return levels
	}
	newLevels := make([]string, 0, len(levels)-1)
	newLevels = append(newLevels, levels[:index]...)
	return append(newLevels, levels[index+1:]...)

}

func main() {
	data := ReadFile("input.txt")

	safeReportsProblem1 := 0
	safeReportsProblem2 := 0

	for _, line := range data {
		levels := strings.Split(line, " ")
		if ReportIsSafe(levels) {
			safeReportsProblem1++
		}

		for index := range levels {
			newLevels := removeAtIndex(levels, index)
			if ReportIsSafe(newLevels) {
				safeReportsProblem2++
				break
			}
		}

	}

	fmt.Println("Problem 1: ", safeReportsProblem1)
	fmt.Println("Problem 2: ", safeReportsProblem2)
}

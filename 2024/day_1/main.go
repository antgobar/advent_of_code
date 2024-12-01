package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func ReadTextFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", path))
	}
	return strings.Split(string(content), "\n")
}

func occurrenceCount(num int, nums []int) int {
	occurrences := 0

	for _, n := range nums {
		if num == n {
			occurrences++
		}
	}
	return occurrences
}

func parseAndSort(data []string) ([]int, []int) {
	var colA []int
	var colB []int

	for _, line := range data {
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)

		numColA, errA := strconv.Atoi(matches[0])
		numColB, errB := strconv.Atoi(matches[1])

		if errA != nil || errB != nil {
			panic(fmt.Sprintf("Error converting numbers %v", line))
		}

		colA = append(colA, numColA)
		colB = append(colB, numColB)
	}

	sort.Ints(colA)
	sort.Ints(colB)
	return colA, colB
}

func main() {
	data := ReadTextFile("input.txt")

	colA, colB := parseAndSort(data)

	totalDistance := 0
	similarityScore := 0

	for i, val := range colA {
		distance := math.Abs(float64(val - colB[i]))
		totalDistance += int(distance)
		
		occurrence := occurrenceCount(val, colB)
		similarityScore += val * occurrence

	}

	fmt.Println("Problem 1:", totalDistance)
	fmt.Println("Problem 2:", similarityScore)
}

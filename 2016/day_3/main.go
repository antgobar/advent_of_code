package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type SidesStr []string

func (strings SidesStr) ToInts() []int {
	ints := make([]int, len(strings))

	for i, s := range strings {
		ints[i], _ = strconv.Atoi(s)
	}
	return ints
}

func ReadTextFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	return strings.Split(string(content), "\n")
}

func main() {
	data := ReadTextFile("day_3.txt")

	count := 0

	for _, line := range data {
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)
		sides := SidesStr(matches).ToInts()
		sort.Ints(sides)

		if sides[0]+sides[1] > sides[2] {
			count += 1
			continue
		}
	}
	fmt.Println(count)
}

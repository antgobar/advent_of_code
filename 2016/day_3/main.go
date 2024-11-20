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

func Problem1(data []string) {
	count := 0

	for _, line := range data {
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(line, -1)
		sides := SidesStr(matches).ToInts()
		sort.Ints(sides)

		if sides[0]+sides[1] > sides[2] {
			count++
		}
	}
	fmt.Println("Problem 1:", count)
}

func Problem2(data []string) {
	count := 0

	for i := 0; i <= len(data); i += 3 {
		if i+3 > len(data) {
			continue
		}
		numb := strings.Join(data[i:i+3], " ")
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(numb, -1)
		sides := SidesStr(matches).ToInts()
		groupA := []int{sides[0], sides[3], sides[6]}
		groupB := []int{sides[1], sides[4], sides[7]}
		groupC := []int{sides[2], sides[5], sides[8]}
		sort.Ints(groupA)
		sort.Ints(groupB)
		sort.Ints(groupC)

		if groupA[0]+groupA[1] > groupA[2] {
			count++
		}
		if groupB[0]+groupB[1] > groupB[2] {
			count++
		}
		if groupC[0]+groupC[1] > groupC[2] {
			count++

		}

	}

	fmt.Println("Problem 2:", count)
}

func main() {
	data := ReadTextFile("day_3.txt")

	Problem1(data)
	Problem2(data)

}

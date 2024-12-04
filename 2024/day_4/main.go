package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) [][]string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file %v", path))
	}

	lines := strings.Split(string(content), "\n")

	var data [][]string

	for _, line := range lines {
		chars := strings.Split(line, "")
		data = append(data, chars)
	}
	return data
}

func getWordsInDirections(array [][]string, x, y, size int) []string {
	var northWest []string
	var north []string
	var northEast []string
	var east []string
	var southEast []string
	var south []string
	var southWest []string
	var west []string

	hasSpaceNorth := y >= size-1
	hasSpaceEast := x < len(array[y])-size
	hasSpaceSouth := y < len(array)-size
	hasSpaceWest := x >= size-1

	for i := 0; i < size; i++ {
		if hasSpaceNorth && hasSpaceWest {
			northWest = append(northWest, array[y-i][x-i])
		}
		if hasSpaceNorth {
			north = append(north, array[y-i][x])
		}
		if hasSpaceNorth && hasSpaceEast {
			northEast = append(northEast, array[y-i][x+i])
		}
		if hasSpaceEast {
			east = append(east, array[y][x+i])
		}
		if hasSpaceSouth && hasSpaceEast {
			southEast = append(southEast, array[y+i][x+i])
		}
		if hasSpaceSouth {
			south = append(south, array[y+i][x])
		}
		if hasSpaceSouth && hasSpaceWest {
			southWest = append(southWest, array[y+i][x-i])
		}
		if hasSpaceWest {
			west = append(west, array[y][x-i])
		}

	}

	directions := [][]string{northWest, north, northEast, east, southEast, south, southWest, west}

	var words []string

	for _, direction := range directions {
		words = append(words, strings.Join(direction, ""))
	}
	return words
}

func validWordsCount(words []string, searchWord string) int {
	count := 0
	for _, word := range words {
		if word == searchWord {
			count++
		}
	}
	return count
}

func main() {
	data := ReadFile("test2.txt")
	const searchWord = "XMAS"
	totalCount := 0
	for y := range data {
		for x := range data[y] {
			words := getWordsInDirections(data, x, y, len(searchWord))
			totalCount += validWordsCount(words, searchWord)
		}
	}

	fmt.Println("Problem 1: ", totalCount)

}

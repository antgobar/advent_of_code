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

func getValidWordsInDirections(array [][]string, x, y int, searchWord string) (int, int) {
	size := len(searchWord)

	var northWest []string
	var north []string
	var northEast []string
	var east []string
	var southEast []string
	var south []string
	var southWest []string
	var west []string

	hasSpaceNorth := y >= size-1
	hasSpaceEast := x < len(array[y])-size+1
	hasSpaceSouth := y < len(array)-size+1
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
	orthogonals := [][]string{north, east, west, south}
	diagonals := [][]string{northWest, northEast, southEast, southWest}
	directions := append(orthogonals, diagonals...)
	count1 := 0
	for _, direction := range directions {
		word := strings.Join(direction, "")
		if word == searchWord {
			count1++
		}
	}
	size = 2
	hasSpaceNorth = y >= size-1
	hasSpaceEast = x < len(array[y])-size+1
	hasSpaceSouth = y < len(array)-size+1
	hasSpaceWest = x >= size-1

	count2 := 0
	if array[y][x] == "A" && hasSpaceNorth && hasSpaceEast && hasSpaceSouth && hasSpaceWest {
		if diagonalMatch(array, x, y, "M", "S", "M", "S") {
			count2++
		}
		if diagonalMatch(array, x, y, "S", "M", "M", "S") {
			count2++
		}
		if diagonalMatch(array, x, y, "M", "S", "S", "M") {
			count2++
		}
		if diagonalMatch(array, x, y, "S", "M", "S", "M") {
			count2++
		}
	}

	return count1, count2
}

func diagonalMatch(array [][]string, x, y int, nw, se, ne, sw string) bool {
	if string(array[y-1][x-1]) == nw && string(array[y+1][x+1]) == se &&
		string(array[y-1][x+1]) == ne && string(array[y+1][x-1]) == sw {
		return true
	}
	return false
}

func main() {
	data := ReadFile("input.txt")
	const searchWord = "XMAS"
	totalCount1 := 0
	totalCount2 := 0
	for y := range data {
		for x := range data[y] {
			count1, count2 := getValidWordsInDirections(data, x, y, searchWord)
			totalCount1 += count1
			totalCount2 += count2
		}
	}

	fmt.Println("Problem 1: ", totalCount1)
	fmt.Println("Problem 2: ", totalCount2)
}

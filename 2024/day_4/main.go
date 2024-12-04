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

type Grid [][]Coordinate

func (g Grid) Size() (int, int) {
	return len(g[0]), len(g)
}

// func (g Grid) getElemsFromCoords(n int) []Coordinate {

// }

type Coordinate struct {
	x, y  int
	value string
}

func NewGrid(data [][]string) Grid {
	sizeY := len(data)
	sizeX := len(data[0])

	coordinates := make([][]Coordinate, sizeY)

	for y := range data {
		row := make([]Coordinate, sizeX)
		for x := range data[y] {
			row[x] = Coordinate{x, y, data[y][x]}
		}
		coordinates[y] = row
	}
	return Grid(coordinates)
}

func hasAheadSpace(array []string, index int, word string) bool {
	if index <= len(array)-len(word) {
		return true
	}
	return false
}

func hasRearSpace(index int, word string) bool {
	if index > len(word)-2 {
		return true
	}
	return false
}

func getDirections(array [][]string, x, y, size int) []string {
	var northWest []string
	var north []string
	var northEast []string
	var east []string
	var southEast []string
	var south []string
	var southWest []string
	var west []string

	for i := 0; i < size; i++ {
		northWest = append(northWest, array[y-i][x-i])
		north = append(north, array[y-i][x])
		northEast = append(northEast, array[y-i][x+i])
		east = append(east, array[y][x+i])
		southEast = append(southEast, array[y+i][x+i])
		south = append(south, array[y+i][x])
		southWest = append(southWest, array[y+i][x-i])
		west = append(west, array[y][x-i])
	}

	directions := [][]string{northWest, north, northEast, east, southEast, south, southWest, west}

	var words []string

	for _, direction := range directions {
		words = append(words, strings.Join(direction, ""))
	}
	return words
}

func east(array [][]string, x, y, size int) string {
	var chars []string
	for i := 0; i < size; i++ {
		chars = append(chars, array[y][x+i])
	}
	return strings.Join(chars, "")
}

func west(array [][]string, x, y, size int) string {
	var chars []string
	for i := 0; i < size; i++ {
		chars = append(chars, array[y][x-i])
	}
	return strings.Join(chars, "")
}

func north(array [][]string, x, y, size int) string {
	var chars []string
	for i := 0; i < size; i++ {
		chars = append(chars, array[y-i][x])
	}
	return strings.Join(chars, "")
}

func south(array [][]string, x, y, size int) string {
	var chars []string
	for i := 0; i < size; i++ {
		chars = append(chars, array[y+i][x])
	}
	return strings.Join(chars, "")
}

func southEast(array [][]string, x, y, size int) string {
	var chars []string
	for i := 0; i < size; i++ {
		chars = append(chars, array[y+i][x+i])
	}
	return strings.Join(chars, "")
}

func northEast(array [][]string, x, y, size int) string {
	var chars []string
	for i := 0; i < size; i++ {
		chars = append(chars, array[y-i][x+i])
	}
	return strings.Join(chars, "")
}

func northWest(array [][]string, x, y, size int) string {
	var chars []string
	for i := 0; i < size; i++ {
		chars = append(chars, array[y-i][x-i])
	}
	return strings.Join(chars, "")
}

func southWest(array [][]string, x, y, size int) string {
	var chars []string
	for i := 0; i < size; i++ {
		chars = append(chars, array[y+i][x-i])
	}
	return strings.Join(chars, "")
}

func main() {
	data := ReadFile("test2.txt")
	const searchWord = "XMAS"
	// for y := range data {
	// 	for x := range data[y] {
	// 		fmt.Println(x, y)
	// 	}
	// }
	fmt.Println(northWest(data, 3, 3, len(searchWord)))
	fmt.Println(north(data, 3, 3, len(searchWord)))
	fmt.Println(northEast(data, 3, 3, len(searchWord)))
	fmt.Println(east(data, 3, 3, len(searchWord)))
	fmt.Println(southEast(data, 3, 3, len(searchWord)))
	fmt.Println(south(data, 3, 3, len(searchWord)))
	fmt.Println(southWest(data, 3, 3, len(searchWord)))
	fmt.Println(west(data, 3, 3, len(searchWord)))

	fmt.Println(getDirections(data, 3, 3, len(searchWord)))
}

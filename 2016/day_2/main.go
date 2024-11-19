package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadTextFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	return strings.Split(string(content), "\n")
}

type Coordinates struct {
	x, y int
}

func (c Coordinates) Movement(direction string) Coordinates {
	var newCoords Coordinates
	switch direction {
	case "U":
		newCoords = Coordinates{c.x, c.y + 1}
	case "D":
		newCoords = Coordinates{c.x, c.y - 1}
	case "L":
		newCoords = Coordinates{c.x - 1, c.y}
	case "R":
		newCoords = Coordinates{c.x + 1, c.y}
	default:
		newCoords = c
	}
	if newCoords.x < 0 {
		return Coordinates{0, newCoords.y}
	}
	if newCoords.x > 2 {
		return Coordinates{2, newCoords.y}
	}
	if newCoords.y < 0 {
		return Coordinates{newCoords.x, 0}
	}
	if newCoords.y > 2 {
		return Coordinates{newCoords.x, 2}
	}
	return newCoords
}

func (c Coordinates) AsDigit() int {
	switch c {
	case Coordinates{0, 0}:
		return 7
	case Coordinates{0, 1}:
		return 8
	case Coordinates{0, 2}:
		return 9
	case Coordinates{1, 0}:
		return 4
	case Coordinates{1, 1}:
		return 5
	case Coordinates{1, 2}:
		return 6
	case Coordinates{2, 0}:
		return 1
	case Coordinates{2, 1}:
		return 2
	case Coordinates{2, 2}:
		return 3
	default:
		return 0
	}
}

func getDigit(line string) int {
	actions := strings.Split(line, "")
	var coordinate = Coordinates{1, 1}

	for _, action := range actions {
		coordinate = coordinate.Movement(action)
	}
	return coordinate.AsDigit()
}

type Digits []int

func (d Digits) Combine() string {
	stringNumbers := make([]string, len(d))
	for i, num := range d {
		stringNumbers[i] = strconv.Itoa(num)
	}
	return strings.Join(stringNumbers, "")
}

func main() {
	data := ReadTextFile("day_2_test.txt")
	var digits Digits

	for _, line := range data {
		digit := getDigit(line)
		digits = append(digits, digit)
	}

	fmt.Println("Problem 1:", digits.Combine())
}

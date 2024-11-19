package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadTextFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	return strings.Split(string(content), "\n")
}

func (c Coordinates) Movement(direction string, numPad NumPad) Coordinates {
	var newCoords Coordinates
	bounds := numPad.Bounds()
	fmt.Println(bounds)
	switch direction {
	case "U":
		newCoords = Coordinates{c.x, c.y + 1}
		if newCoords.y > bounds.yMax {
			return Coordinates{newCoords.x, bounds.yMax}
		}
		if numPad.ValueAtCoord(newCoords) == " " {
			return c
		}
		return newCoords
	case "D":
		newCoords = Coordinates{c.x, c.y - 1}
		if newCoords.y < bounds.yMin {
			return Coordinates{newCoords.x, bounds.yMin}
		}
		if numPad.ValueAtCoord(newCoords) == " " {
			return c
		}
		return newCoords
	case "L":
		newCoords = Coordinates{c.x - 1, c.y}
		if newCoords.x < bounds.xMin {
			return Coordinates{bounds.xMin, newCoords.y}
		}
		if numPad.ValueAtCoord(newCoords) == " " {
			return c
		}
		return newCoords
	case "R":
		newCoords = Coordinates{c.x + 1, c.y}
		if newCoords.x > bounds.xMax {
			return Coordinates{bounds.xMax, newCoords.y}
		}
		if numPad.ValueAtCoord(newCoords) == " " {
			return c
		}
		return newCoords
	default:
		return c
	}
}

func getDigit(line string, coordinate Coordinates, numPad NumPad) Coordinates {
	actions := strings.Split(line, "")

	for _, action := range actions {
		coordinate = coordinate.Movement(action, numPad)
	}
	return coordinate
}

type Digits []string

func main() {
	data := ReadTextFile("day_2.txt")

	var digits1 Digits
	var startCoordinate1 = Coordinates{1, 1}
	numberPad1 := NumPad{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	var digits2 Digits
	var startCoordinate2 = Coordinates{0, 2}
	numberPad2 := NumPad{
		{" ", " ", "1", " ", " "},
		{" ", "2", "3", "4", " "},
		{"5", "6", "7", "8", "9"},
		{" ", "A", "B", "C", " "},
		{" ", " ", "D", " ", " "},
	}

	for _, line := range data {
		coordinate1 := getDigit(line, startCoordinate1, numberPad1)
		startCoordinate1 = coordinate1
		digits1 = append(digits1, numberPad1.ValueAtCoord(coordinate1))

		coordinate2 := getDigit(line, startCoordinate2, numberPad2)
		startCoordinate2 = coordinate2
		digits2 = append(digits2, numberPad2.ValueAtCoord(coordinate2))
	}

	fmt.Println("Problem 1:", strings.Join(digits1, ""))
	fmt.Println("Problem 2:", strings.Join(digits2, ""))

}

type Coordinates struct {
	x, y int
}

type NumPad [][]string

type NumPadBounds struct {
	xMax, xMin, yMax, yMin int
}

func (np NumPad) ValueAtCoord(c Coordinates) string {
	x := max(0, min(c.x, len(np[0])-1))
	y := max(0, min(len(np)-1-c.y, len(np)-1))
	return np[y][x]
}

func (np NumPad) Bounds() NumPadBounds {
	return NumPadBounds{
		xMax: len(np[0]) - 1,
		yMax: len(np) - 1,
	}
}

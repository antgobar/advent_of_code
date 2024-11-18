package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func rotationToDirection(direction string) int {
	return map[string]int{
		"L": -1, "R": 1,
	}[direction]
}

func main() {
	content, err := os.ReadFile("day_1.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	instructions := strings.Split(string(content), ", ")

	var visitedCoordinates coordinateContainer
	var angle float64 = 0
	currentCoordinates := coordinates{0, 0}

	visitedCoordinates = append(visitedCoordinates, currentCoordinates)

	for _, instruction := range instructions {
		var rotation = string(instruction[0])
		steps, err := strconv.Atoi(instruction[1:])
		if err != nil {
			fmt.Println("Error converting steps:", err)
			return
		}

		direction := rotationToDirection(rotation)

		rotationAngle := float64(direction) * 90.0 * (math.Pi / 180)
		angle += rotationAngle
		sineValue := int(math.Sin(angle))   // x dir
		cosineValue := int(math.Cos(angle)) // y dir
		xStep := steps * sineValue
		yStep := steps * cosineValue
		if xStep != 0 && yStep == 0 {
			for i := 0; i < int(math.Abs(float64(xStep))); i++ {
				currentCoordinates = coordinates{currentCoordinates.x + sineValue, currentCoordinates.y + 0}
				visitedCoordinates = append(visitedCoordinates, currentCoordinates)
			}
		}

		if xStep == 0 && yStep != 0 {
			for i := 0; i < int(math.Abs(float64(yStep))); i++ {
				currentCoordinates = coordinates{currentCoordinates.x + 0, currentCoordinates.y + cosineValue}
				visitedCoordinates = append(visitedCoordinates, currentCoordinates)
			}

		}

	}
	fmt.Println("Problem 1:", currentCoordinates.ShortestPath())
	fmt.Println("Problem 2:", visitedCoordinates.findFirstDuplicate().ShortestPath())
}

func (container coordinateContainer) findFirstDuplicate() coordinates {
	seen := make(map[coordinates]bool)

	for _, item := range container {
		if seen[item] {
			return item
		}
		seen[item] = true
	}
	return coordinates{0, 0}
}

type coordinates struct {
	x, y int
}

func (c coordinates) ShortestPath() int {
	return int(math.Abs(float64(c.x)) + math.Abs(float64(c.y)))
}

type coordinateContainer []coordinates

func (container coordinateContainer) Contains(coords coordinates) bool {
	for _, storedCoords := range container {
		if storedCoords.x == coords.x && storedCoords.y == coords.x {
			return true
		}
	}
	return false
}

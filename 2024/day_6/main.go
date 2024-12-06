package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) Grid {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file %v", path))
	}

	var grid Grid

	rows := strings.Split(string(content), "\n")

	for _, row := range rows {
		var col []string
		items := strings.Split(row, "")
		for _, item := range items {
			col = append(col, item)
		}
		grid = append(grid, items)
	}
	return grid
}

type Grid [][]string

type Boundary struct {
	x, y int
}

func (g Grid) Size() Boundary {
	return Boundary{len(g[0]), len(g)}
}

type Coordinate struct {
	x, y int
}

func (g Grid) StartingPoint() Coordinate {
	for y := range g {
		for x := range g[y] {
			if g[y][x] == "^" {
				return Coordinate{x, y}
			}
		}
	}
	return Coordinate{}
}

type Direction string

type Guard struct {
	grid      Grid
	position  Coordinate
	direction Direction
}

func (g Guard) IsInGrid() bool {
	gx, gy := g.position.x, g.position.y

	if gy < 0 || gy > g.grid.Size().y-1 || gx < 0 || gx > g.grid.Size().x-1 {
		return false
	}
	return true
}

func (guard Guard) NextStep(grid Grid) Guard {
	return guard
}

func main() {
	grid := ReadFile("test.txt")
	guard := Guard{grid, grid.StartingPoint(), "^"}
	fmt.Print(guard.IsInGrid())
}

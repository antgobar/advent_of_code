package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", path))
	}

	return strings.Split(string(content), "\n")
}

func main() {
	lines := ReadFile("test.txt")

	var arr [][]string

	for _, line := range lines {
		arr = append(arr, strings.Split(line, ""))
	}

	fmt.Println(arr[1][2])

}

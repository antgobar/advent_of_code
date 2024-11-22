package main

import (
	"fmt"
	"os"
	"sort"
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

type Room struct {
	ID       int
	Name     string
	CheckSum string
}

func ParseLineToRoom(line string) Room {
	rawRoom := strings.Split(line, "[")
	preRoomNameId, preCheckSum := rawRoom[0], rawRoom[1]
	splitRoomNameId := strings.Split(preRoomNameId, "-")

	roomIdStr := splitRoomNameId[len(splitRoomNameId)-1]
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		fmt.Println("Error:", err)
	}
	roomName := strings.Join(splitRoomNameId[:len(splitRoomNameId)-1], "-")
	checkSum := preCheckSum[:len(preCheckSum)-1]

	return Room{ID: roomId, Name: roomName, CheckSum: checkSum}
}

type Frequency struct {
	char  rune
	count int
}

func (r Room) GenerateCheckSum() string {
	char_frequencies := make(map[rune]int)
	var frequencies []Frequency

	roomName := strings.ReplaceAll(r.Name, "-", "")

	for _, char := range roomName {
		char_frequencies[char]++
	}

	for char, count := range char_frequencies {
		frequencies = append(frequencies, Frequency{char, count})
	}

	sort.Slice(frequencies, func(i, j int) bool {
		if frequencies[i].count != frequencies[j].count {
			return frequencies[i].count > frequencies[j].count
		}
		return frequencies[i].char < frequencies[j].char
	})

	result := make([]rune, 0, 5)
	for i := 0; i < 5 && i < len(frequencies); i++ {
		result = append(result, frequencies[i].char)
	}

	return string(result)
}

func (r Room) IsReal() bool {
	return r.GenerateCheckSum() == r.CheckSum
}

func (r Room) RotN() string {
	result := make([]rune, 0, len(r.Name))

	for _, char := range r.Name {
		var rotated rune
		if char >= 'A' && char <= 'Z' {
			rotated = 'A' + (char-'A'+rune(r.ID))%26
		} else if char >= 'a' && char <= 'z' {
			rotated = 'a' + (char-'a'+rune(r.ID))%26
		} else {
			rotated = char
		}

		result = append(result, rotated)
	}

	return string(result)
}

func main() {
	data := ReadTextFile("day_4.txt")
	count := 0
	for _, line := range data {
		room := ParseLineToRoom(line)
		if room.RotN() == "northpole-object-storage" {
			fmt.Println("Problem 2:", room.ID)
		}
		if room.IsReal() {
			count += room.ID
		}
	}

	fmt.Println("Problem 1:", count)
}

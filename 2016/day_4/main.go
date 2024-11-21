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
	roomName := strings.Join(splitRoomNameId[:len(splitRoomNameId)-1], "")
	checkSum := preCheckSum[:len(preCheckSum)-1]

	return Room{ID: roomId, Name: roomName, CheckSum: checkSum}
}

func (r Room) GenerateCheckSum() string {
	for _, char := range strings.Split(r.Name, "") {
		fmt.Println(char)
	}

	return "foo"
}

func (r Room) IsReal() bool {
	return r.GenerateCheckSum() == r.CheckSum
}

func main() {
	data := ReadTextFile("day_4_test.txt")
	count := 0
	for _, line := range data {
		room := ParseLineToRoom(line)
		if room.IsReal() {
			count += room.ID
		}
	}

	fmt.Println("Problem 1:", count)
}

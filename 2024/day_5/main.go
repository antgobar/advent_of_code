package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func print(v ...interface{}) {
	fmt.Println(v...)
}

func MustStr2Int(str string) int {
	val, err := strconv.Atoi(str)
	if err != err {
		panic(fmt.Sprintf("Error converting string %v to int", str))
	}
	return val
}

func MustStr2IntArray(strs []string) []int {
	var nums []int
	for _, str := range strs {
		nums = append(nums, MustStr2Int(str))
	}
	return nums
}

type Rule struct {
	first, second int
}

type Rules []Rule

type Updates []Update

type Update []int

func (rule Rule) IsFollowed(update Update) bool {
	first := update.GetIndexAtValue(rule.first)
	second := update.GetIndexAtValue(rule.second)
	if first == -1 || second == -1 {
		return true
	}
	if second > first {
		return true
	}
	return false
}

func (update Update) GetIndexAtValue(val int) int {
	for i, num := range update {
		if num == val {
			return i
		}
	}
	return -1
}

func (update Update) FollowsAllRules(rules Rules) bool {
	for _, rule := range rules {
		if !rule.IsFollowed(update) {
			return false
		}
	}
	return true
}

func (u Update) Middle() int {
	return u[len(u)/2]
}

func parseRules(rawRules string) Rules {
	rules := strings.Split(rawRules, "\n")

	var parsedRules Rules

	for _, rule := range rules {
		firstSecond := strings.Split(rule, "|")
		parsedRules = append(parsedRules, Rule{MustStr2Int(firstSecond[0]), MustStr2Int(firstSecond[1])})
	}
	return parsedRules
}

func parseUpdates(rawUpdates string) Updates {
	updates := strings.Split(rawUpdates, "\n")

	var parsedUpdates Updates

	for _, update := range updates {
		nums := MustStr2IntArray(strings.Split(update, ","))
		parsedUpdates = append(parsedUpdates, nums)
	}
	return parsedUpdates
}

func ReadFile(path string) (Rules, Updates) {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Sprintf("Error reading file %v", path))
	}

	data := strings.Split(string(content), "\n\n")
	rules := parseRules(data[0])
	updates := parseUpdates(data[1])
	return rules, updates
}

func main() {
	rules, updates := ReadFile("input.txt")
	sumProblem1 := 0
	for _, update := range updates {
		if update.FollowsAllRules(rules) {
			sumProblem1 += update.Middle()
		}
	}
	print("Problem 1:", sumProblem1)
}

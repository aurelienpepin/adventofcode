package day6

import (
	"io/ioutil"
	"log"
	"strings"
)

func common() []string {
	content, err := ioutil.ReadFile("2020/inputs/day6")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n\n")
}

func Part1() int {
	groups := common()
	total := 0

	for _, group := range groups {
		total += countYes(group)
	}

	return total
}

// Using Vanilla Go
func Part2() int {
	groups := common()
	total := 0

	for _, group := range groups {
		total += countYesFromEveryone(group)
	}

	return total
}

func countYes(group string) int {
	questionSet := make(map[int32]bool)

	for _, char := range group {
		if 'a' <= char && char <= 'z' {
			questionSet[char] = true
		}
	}

	return len(questionSet)
}

func countYesFromEveryone(group string) int {
	people := strings.Split(group, "\n")
	questions := make([]int, 26)

	for _, person := range people {
		for _, char := range person {
			questions[char - 'a']++
		}
	}

	// Count questions answered by everyone
	totalYes := 0

	for _, count := range questions {
		if count == len(people) {
			totalYes++
		}
	}

	return totalYes
}
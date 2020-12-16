package day16

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type fieldInterval struct {
	name	string
	min1	int
	max1	int
	min2	int
	max2	int
}

func (interval fieldInterval) contains(number int) bool {
	return (interval.min1 <= number && number <= interval.max1) || (interval.min2 <= number && number <= interval.max2)
}

func common() ([]fieldInterval, []int, [][]int) {
	content, err := ioutil.ReadFile("2020/inputs/day16")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(content), "\n\n")
	return toIntervals(parts[0]), toYourTicket(strings.Split(parts[1], "\n")[1]), toNearbyTickets(parts[2])
}

/**
	PART 1
 */
func Part1() int {
	intervals, _, nearbyTickets := common()
	sum := 0

	for _, ticket := range nearbyTickets {
		for _, value := range ticket {
			if !isValidFor(intervals, value) {
				sum += value
			}
		}
	}

	return sum
}

func isValidFor(intervals []fieldInterval, number int) bool {
	for _, interval := range intervals {
		if interval.contains(number) {
			return true
		}
	}

	return false
}

/**
	PARSING FUNCTIONS
 */
func toIntervals(input string) []fieldInterval {
	lines := strings.Split(input, "\n")
	var intervals []fieldInterval

	for _, line := range lines {
		intervals = append(intervals, toInterval(line))
	}
	return intervals
}

func toInterval(input string) fieldInterval {
	parts := strings.Split(input, ": ")
	name := parts[0]

	subparts := strings.Split(parts[1], " or ")
	min1, max1 := toRange(subparts[0])
	min2, max2 := toRange(subparts[1])

	return fieldInterval{name, min1, max1, min2, max2}
}

func toRange(input string) (int, int) {
	parts := strings.Split(input, "-")
	return convOrPanic(parts[0]), convOrPanic(parts[1])
}

func toYourTicket(input string) []int {
	parts := strings.Split(input, ",")
	var numbers []int

	for _, part := range parts {
		numbers = append(numbers, convOrPanic(part))
	}
	return numbers
}

func toNearbyTickets(input string) [][]int {
	lines := strings.Split(input, "\n")[1:]
	var tickets [][]int

	for _, line := range lines {
		tickets = append(tickets, toYourTicket(line))
	}
	return tickets
}

func convOrPanic(number string) int {
	if value, err := strconv.Atoi(number); err == nil {
		return value
	} else {
		panic("error, integer impossible to cast: " + number)
	}
}
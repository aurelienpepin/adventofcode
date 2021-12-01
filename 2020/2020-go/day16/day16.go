package day16

import (
	"io/ioutil"
	"log"
	"sort"
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

type matching struct {
	index		int
	intervals	[]fieldInterval
}

type byLength []matching

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
	PART 3
 */
func Part2() int {
	intervals, ticket, tickets := common()
	tickets = discardInvalid(tickets, intervals)
	tickets = append(tickets, ticket)

	ticketProduct := 1
	var possibilities []matching
	for i := 0; i < len(tickets[0]); i++ {
		possibilities = append(possibilities, matching{i, listPossibilitiesForColumn(i, tickets, intervals)})
	}

	for len(possibilities) > 0 {
		i, col, newPossibilities := findOneGoodColumn(possibilities)
		possibilities = newPossibilities

		if strings.HasPrefix(col.name, "departure") {
			ticketProduct *= ticket[i]
		}
	}

	return ticketProduct
}

func listPossibilitiesForColumn(index int, tickets [][]int, intervals []fieldInterval) []fieldInterval {
	var candidates []fieldInterval

	checkNextInterval:
	for _, interval := range intervals {
		for _, ticket := range tickets {
			if !interval.contains(ticket[index]) {
				continue checkNextInterval
			}
		}

		candidates = append(candidates, interval)
	}

	return candidates
}

func findOneGoodColumn(matchings []matching) (int, fieldInterval, []matching) {
	sort.Sort(byLength(matchings))
	trivialMatching := matchings[0]

	// Remove the new selected interval from other matchings
	for i := 0; i < len(matchings); i++ {
		matchings[i].remove(trivialMatching.intervals[0])
	}

	// Remove the empty matching
	matchings[0] = matchings[len(matchings) - 1]
	return trivialMatching.index, trivialMatching.intervals[0], matchings[:len(matchings) - 1]
}

func (matching *matching) remove(interval fieldInterval) {
	for i := 0; i < len(matching.intervals); i++ {
		if matching.intervals[i].name == interval.name {
			matching.intervals[i] = matching.intervals[len(matching.intervals) - 1]
			break
		}
	}
	matching.intervals = matching.intervals[:len(matching.intervals) - 1]
}

func discardInvalid(tickets [][]int, intervals []fieldInterval) [][]int {
	var validTickets [][]int

	checkNewTicket:
	for _, ticket := range tickets {
		for _, value := range ticket {
			if !isValidFor(intervals, value) {
				continue checkNewTicket
			}
		}
		validTickets = append(validTickets, ticket)
	}
	return validTickets
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

/**
	SORT WITH A CUSTOM FUNCTION
	The sorted type (an alias for the real type) must implement three functions:
 */
func (b byLength) Len() int {
	return len(b)
}

func (b byLength) Less(i, j int) bool {
	return len(b[i].intervals) < len(b[j].intervals)
}

func (b byLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

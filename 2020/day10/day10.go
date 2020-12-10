package day10

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func common() []int {
	content, err := ioutil.ReadFile("2020/inputs/day10")
	if err != nil {
		log.Fatal(err)
	}

	var numbers []int
	for _, line := range strings.Split(string(content), "\n") {
		if number, err := strconv.Atoi(line); err == nil {
			numbers = append(numbers, number)
		} else {
			log.Fatal(err)
		}
	}

	return numbers
}

func Part1() int {
	adapters := common()
	sort.Ints(adapters)

	differences := make(map[int]int)
	differences[adapters[0]]++

	for i := 1; i < len(adapters); i++ {
		differences[adapters[i] - adapters[i - 1]]++
	}

	differences[3]++
	return differences[1] * differences[3]
}

func Part2() uint64 {
	adapters := common()
	sort.Ints(adapters)

	join := make([]uint64, len(adapters))
	join[0] = 1

	for i := 1; i < len(adapters); i++ {
		if adapters[i] <= 3 {
			join[i]++
		}

		for j := i - 1; j >= 0 && adapters[i] - adapters[j] <= 3; j-- {
			join[i] += join[j]
		}
	}

	return join[len(adapters) - 1]
}
package day9

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func common() []int {
	content, err := ioutil.ReadFile("2020/inputs/day9")
	if err != nil {
		log.Fatal(err)
	}

	var input []int
	for _, line := range strings.Split(string(content), "\n") {
		if number, err := strconv.Atoi(line); err == nil {
			input = append(input, number)
		}
	}

	return input
}

func Part1() int {
	numbers := common()
	window := 25

	sums := make(map[int]int)

	// Init the sums map
	for i := 0; i < window; i++ {
		for j := i + 1; j < window; j++ {
			sums[numbers[i] + numbers[j]]++
		}
	}

	// fmt.Println(sums)

	for i := window; i < len(numbers); i++ {
		if v, ok := sums[numbers[i]]; !ok || v == 0 {
			return numbers[i] // bad number
		}

		toRemove := numbers[i - window]
		toAdd := numbers[i]

		for _, elem := range numbers[i - window + 1:i] {
			sums[elem + toRemove]--
			sums[elem + toAdd]++
		}

		// fmt.Println(toRemove, toAdd, sums)
	}

	panic("bad number does not exist")
}

func Part2() int {
	invalidNumber := Part1()
	fmt.Println(invalidNumber)

	numbers := common()

	// Naive solution O(n^2)
	for i := 0; i < len(numbers); i++ {
		partialSum := numbers[i]
		min, max := numbers[i], numbers[i]

		for j := i + 1; j < len(numbers); j++ {
			partialSum += numbers[j]
			if numbers[j] < min {
				min = numbers[j]
			} else if numbers[j] > max {
				max = numbers[j]
			}

			if partialSum == invalidNumber {
				return min + max
			} else if partialSum > invalidNumber {
				break
			}
		}
	}

	panic("subarray not found")
}
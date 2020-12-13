package day13

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func common() (int, []int) {
	content, err := ioutil.ReadFile("2020/inputs/day13")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(content), "\n")
	departure, err := strconv.Atoi(input[0])
	if err != nil {
		log.Fatal(err)
	}

	var buses []int
	for _, number := range strings.Split(input[1], ",") {
		if bus, err := strconv.Atoi(number); err == nil {
			buses = append(buses, bus)
		} else {
			// ignore 'x'
		}
	}

	return departure, buses
}

func Part1() int {
	departure, buses := common()
	minDelta, minBus := math.MaxInt32, 0

	for _, bus := range buses {
		delta := smallestMultipleOfGreaterThanOrEqualTo(bus, departure) - departure
		if delta < minDelta {
			minDelta = delta
			minBus = bus
		}
	}

	return minDelta * minBus
}

func smallestMultipleOfGreaterThanOrEqualTo(x int, threshold int) int {
	return int(math.Ceil(float64(threshold) / float64(x))) * x
}
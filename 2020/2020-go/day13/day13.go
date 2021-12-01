package day13

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func common(ignore bool) (int, []int) {
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
			if !ignore {
				buses = append(buses, -1)
			}
		}
	}

	return departure, buses
}

func Part1() int {
	departure, buses := common(true)
	minDelta, minBus := math.MaxInt64, 0

	for _, bus := range buses {
		delta := smallestMultipleOfGreaterThanOrEqualTo(bus, departure) - departure
		if delta < minDelta {
			minDelta = delta
			minBus = bus
		}
	}

	return minDelta * minBus
}

func Part2() int64 {
	_, buses := common(false)

	// This relates to the Chinese remainder theorem!
	var a, n []int64
	for i, bus := range buses {
		if bus > 0 {
			a = append(a, int64(-i))
			n = append(n, int64(bus))
		}
	}

	A, N := chineseRemainderTheorem(a, n)
	for A > N {
		A -= N
	}

	for A < 0 {
		A += N
	}

	return A
}

func smallestMultipleOfGreaterThanOrEqualTo(x int, threshold int) int {
	return int(math.Ceil(float64(threshold) / float64(x))) * x
}

func chineseRemainderTheorem(a []int64, n []int64) (int64, int64) {
	N := int64(1)
	for _, nk := range n {
		N *= nk
	}

	A := int64(0)
	for i := 0; i < len(a); i++ {
		_, _, v := extendedGcd(n[i], N / n[i])
		A += a[i] * v * (N / n[i])
	}

	return A, N
}

func extendedGcd(p int64, q int64) (d int64, a int64, b int64) {
	if q == 0 {
		return p, 1, 0
	}

	v0, v1, v2 := extendedGcd(q, p % q)
	d = v0
	a = v2
	b = v1 - (p / q) * v2
	return d, a, b
}


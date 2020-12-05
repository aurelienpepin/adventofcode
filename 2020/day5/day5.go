package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	file, err := os.Open("2020/inputs/day5")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var seats []string

	for scanner.Scan() {
		seats = append(seats, scanner.Text())
	}

	min, max, sum := math.MaxInt32, math.MinInt32, 0
	for _, seat := range seats {
		seatRank := rank(seat)
		sum += seatRank

		if seatRank > max {
			max = seatRank
		}

		if seatRank < min {
			min = seatRank
		}
	}

	// Part 1
	fmt.Println(max)

	// Part 2
	fmt.Println(findMissing(min, max, sum))
}

func rank(seat string) int {
	return toRow(seat[:7]) * 8 + toColumn(seat[7:])
}

func toRow(seat string) int {
	min, max := 0, 127
	for _, v := range seat {
		if v == 'F' {
			max = int(float64(max+min) / 2)
		} else {
			min = int(math.Ceil(float64(max+min) / 2))
		}
	}

	return min
}

func toColumn(seat string) int {
	min, max := 0, 7
	for _, v := range seat {
		if v == 'L' {
			max = int(float64(max+min) / 2)
		} else {
			min = int(math.Ceil(float64(max+min) / 2))
		}
	}

	return min
}

func sum0ToN(n int) int {
	return n * (n + 1) / 2
}

func findMissing(min int, max int, sum int) int {
	realSum := sum0ToN(max) - sum0ToN(min - 1)
	return realSum - sum
}
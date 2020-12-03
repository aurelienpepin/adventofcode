package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("2020/inputs/day3")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var forest []string

	for scanner.Scan() {
		forest = append(forest, scanner.Text())
	}

	// Part 1
	fmt.Println(part1Traversing(forest, 3, 1))

	// Part 2
	slope1 := part1Traversing(forest, 1, 1)
	slope2 := part1Traversing(forest, 3, 1)
	slope3 := part1Traversing(forest, 5, 1)
	slope4 := part1Traversing(forest, 7, 1)
	slope5 := part1Traversing(forest, 1, 2)
	fmt.Println(slope1 * slope2 * slope3 * slope4 * slope5)
}

func part1Traversing(forest []string, deltaRight int, deltaBottom int) int {
	y := 0
	trees := 0

	for x := 0; x < len(forest); x += deltaBottom {
		if string(forest[x][y]) == "#" {
			trees++
		}

		y = (y + deltaRight) % len(forest[x])
	}

	return trees
}
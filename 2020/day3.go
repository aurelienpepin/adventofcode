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

	fmt.Println(part1Traversing(forest))
}

func part1Traversing(forest []string) int {
	deltaRight := 3
	y := 0
	trees := 0

	for x := 1; x < len(forest); x++ {
		y = (y + deltaRight) % len(forest[x])

		if string(forest[x][y]) == "#" {
			trees++
		}
	}

	return trees
}
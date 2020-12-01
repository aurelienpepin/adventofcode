package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2020/inputs/day1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var numbers []int

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, x)
	}

	// Which integers sum to 2000
	part1(numbers)
	part2(numbers)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part1(numbers []int) {
	numberSet := make(map[int]bool)

	for _, number:= range numbers {
		_, ok := numberSet[2020 - number]
		if ok {
			fmt.Println(number, "and", 2020 - number, "give", number * (2020 - number))
			return
		}

		numberSet[number] = true
	}
}

func part2(numbers []int) {

}

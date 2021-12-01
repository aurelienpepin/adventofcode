package day15

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type gameData struct {
	lastTurn		int
	lastLastTurn	int
}

func common() []int {
	content, err := ioutil.ReadFile("2020/inputs/day15")
	if err != nil {
		log.Fatal(err)
	}

	var numbers []int
	for _, item := range strings.Split(string(content), ",") {
		if number, err := strconv.Atoi(item); err == nil {
			numbers = append(numbers, number)
		} else {
			panic("cannot read number from list")
		}
	}

	return numbers
}

func playTheGameUntil(maxTurn int) int {
	numbers, turn := common(), 1
	stats := make(map[int]gameData)
	lastSpoken := numbers[len(numbers) - 1]

	for _, number := range numbers {
		stats[number] = gameData{turn, -1}
		turn++
	}

	for {
		if wasFirstTime(lastSpoken, stats) {
			lastSpoken = 0
		} else {
			lastSpoken = stats[lastSpoken].lastTurn - stats[lastSpoken].lastLastTurn
		}

		update(lastSpoken, turn, stats)
		if turn == maxTurn {
			return lastSpoken
		}

		//fmt.Println("turn", turn, ":", lastSpoken)
		turn++
	}
}

func Part1() int {
	return playTheGameUntil(2020)
}

func Part2() int {
	return playTheGameUntil(30000000)
}

func wasFirstTime(lastSpoken int, stats map[int]gameData) bool {
	if data, ok := stats[lastSpoken]; ok {
		return data.lastLastTurn == -1
	}

	return false
}

func update(lastSpoken int, turn int, stats map[int]gameData) {
	if data, ok := stats[lastSpoken]; ok {
		stats[lastSpoken] = gameData{turn, data.lastTurn}
	} else {
		stats[lastSpoken] = gameData{turn, -1}
	}
}
package day15

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type GameData struct {
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

func Part1() int {
	numbers, turn := common(), 1
	stats := make(map[int]GameData)
	lastSpoken := numbers[len(numbers) - 1]

	for _, number := range numbers {
		stats[number] = GameData{turn, -1}
		turn++
	}

	for {
		if wasFirstTime(lastSpoken, stats) {
			lastSpoken = 0
		} else {
			lastSpoken = stats[lastSpoken].lastTurn - stats[lastSpoken].lastLastTurn
		}

		update(lastSpoken, turn, stats)
		if turn == 2020 {
			return lastSpoken
		}

		//fmt.Println("turn", turn, ":", lastSpoken)
		turn++
	}
}

func wasFirstTime(lastSpoken int, stats map[int]GameData) bool {
	if data, ok := stats[lastSpoken]; ok {
		return data.lastLastTurn == -1
	} else {
		return false
	}
}

func update(lastSpoken int, turn int, stats map[int]GameData) {
	if _, ok := stats[lastSpoken]; ok {
		stats[lastSpoken] = GameData{turn, stats[lastSpoken].lastTurn}
	} else {
		stats[lastSpoken] = GameData{turn, -1}
	}
}
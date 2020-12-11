package day11

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	FLOOR		= '.'
	OCCUPIED 	= '#'
	AVAILABLE	= 'L'
)

func common() [][]rune {
	content, err := ioutil.ReadFile("2020/inputs/day11")
	if err != nil {
		log.Fatal(err)
	}

	var grid [][]rune
	for _, line := range strings.Split(string(content), "\n") {
		grid = append(grid, []rune(line))
	}

	return grid
}

func Part1() int {
	grid, newGrid := common(), common()
	iterations := 0

	for {
		iterations++
		changed := iterate(grid, newGrid)

		if !changed {
			break
		}

		tmp := grid
		grid = newGrid
		newGrid = tmp
	}

	return countOccupied(newGrid)
}

func countOccupied(grid [][]rune) int {
	occupied := 0
	for _, line := range grid {
		for _, seat := range line {
			if seat == OCCUPIED {
				occupied++
			}
		}
	}

	return occupied
}

func iterate(fromGrid [][]rune, toGrid [][]rune) bool {
	atLeastOneChange := false
	for i := 0; i < len(fromGrid); i++ {
		for j := 0; j < len(fromGrid[i]); j++ {
			newSeat, changed := tweak(fromGrid, i, j)
			toGrid[i][j] = newSeat
			atLeastOneChange = atLeastOneChange || changed
		}
	}
	return atLeastOneChange
}

func tweak(fromGrid [][]rune, i int, j int) (rune, bool) {
	neighbours := countNeighbours(fromGrid, i, j)
	if fromGrid[i][j] == AVAILABLE && neighbours == 0 {
		return OCCUPIED, true
	} else if fromGrid[i][j] == OCCUPIED && neighbours >= 4 {
		return AVAILABLE, true
	} else {
		return fromGrid[i][j], false // doesn't change
	}
}

func countNeighbours(fromGrid [][]rune, i int, j int) int {
	neighbours := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if !(dx == 0 && dy == 0) {
				neighbours += occupied(fromGrid, i+dx, j+dy)
			}
		}
	}

	return neighbours
}

func occupied(fromGrid [][]rune, i int, j int) int {
	if i >= 0 && i < len(fromGrid) && j >= 0 && j < len(fromGrid[i]) && fromGrid[i][j] == OCCUPIED {
		return 1
	} else {
		return 0
	}
}

func printGrid(fromGrid [][]rune) {
	for _, line := range fromGrid {
		for _, seat := range line {
			fmt.Print(string(seat))
		}

		fmt.Println()
	}
}
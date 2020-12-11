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

type tweakFunction func([][]rune, int, int) (rune, bool)

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
		changed := iterate(grid, newGrid, tweak)

		if !changed {
			break
		}

		tmp := grid
		grid = newGrid
		newGrid = tmp
	}

	return countOccupied(newGrid)
}

func Part2() int {
	grid, newGrid := common(), common()
	iterations := 0

	for {
		iterations++
		changed := iterate(grid, newGrid, tweakAgain)

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

func iterate(fromGrid [][]rune, toGrid [][]rune, f tweakFunction) bool {
	atLeastOneChange := false
	for i := 0; i < len(fromGrid); i++ {
		for j := 0; j < len(fromGrid[i]); j++ {
			newSeat, changed := f(fromGrid, i, j)
			toGrid[i][j] = newSeat
			atLeastOneChange = atLeastOneChange || changed
		}
	}
	return atLeastOneChange
}

/*********************
		PART 1
 *********************/

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
	if inGrid(fromGrid, i, j) && fromGrid[i][j] == OCCUPIED {
		return 1
	} else {
		return 0
	}
}

/*********************
		PART 2
 *********************/

func tweakAgain(fromGrid[][] rune, i int, j int) (rune, bool) {
	neighbours := countDistantNeighbours(fromGrid, i, j)
	if fromGrid[i][j] == AVAILABLE && neighbours == 0 {
		return OCCUPIED, true
	} else if fromGrid[i][j] == OCCUPIED && neighbours >= 5 {
		return AVAILABLE, true
	} else {
		return fromGrid[i][j], false // doesn't change
	}
}

func countDistantNeighbours(fromGrid [][]rune, i int, j int) int {
	neighbours := 0
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if !(dx == 0 && dy == 0) {
				neighbours += occupiedDirectional(fromGrid, i, j, dx, dy)
			}
		}
	}

	return neighbours
}

func occupiedDirectional(fromGrid [][]rune, i int, j int, dx int, dy int) int {
	x, y := i + dx, j + dy
	for {
		if !inGrid(fromGrid, x, y) || fromGrid[x][y] == AVAILABLE {
			return 0
		} else if fromGrid[x][y] == OCCUPIED {
			return 1
		} else {
			x += dx
			y += dy
		}
	}
}

func inGrid(fromGrid [][]rune, i int, j int) bool {
	return i >= 0 && i < len(fromGrid) && j >= 0 && j < len(fromGrid[i])
}

func printGrid(fromGrid [][]rune) {
	for _, line := range fromGrid {
		for _, seat := range line {
			fmt.Print(string(seat))
		}

		fmt.Println()
	}
}
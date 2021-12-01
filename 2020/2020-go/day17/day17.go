package day17

import (
	"io/ioutil"
	"log"
	"strings"
)

type coord struct {
	x	int
	y	int
	z	int
}

type coord4d struct {
	x	int
	y	int
	z	int
	w	int
}

func common() (map[coord]rune, []coord) {
	content, err := ioutil.ReadFile("2020/inputs/day17")
	if err != nil {
		log.Fatal(err)
	}

	universe := make(map[coord]rune)
	var activeCubes []coord

	for i, line := range strings.Split(string(content), "\n") {
		for j, _ := range line {
			universe[coord{i, j, 0}] = rune(line[j])
			if rune(line[j]) == '#' {
				activeCubes = append(activeCubes, coord{i, j, 0})
			}
		}
	}
	return universe, activeCubes
}

func common4d() (map[coord4d]rune, []coord4d) {
	content, err := ioutil.ReadFile("2020/inputs/day17")
	if err != nil {
		log.Fatal(err)
	}

	universe := make(map[coord4d]rune)
	var activeCubes []coord4d

	for i, line := range strings.Split(string(content), "\n") {
		for j, _ := range line {
			universe[coord4d{i, j, 0, 0}] = rune(line[j])
			if rune(line[j]) == '#' {
				activeCubes = append(activeCubes, coord4d{i, j, 0, 0})
			}
		}
	}
	return universe, activeCubes
}

/**
	PART 1
 */
func Part1() int {
	universe, activeCubes := common()
	cycleNumber := 6

	for i := 0; i < cycleNumber; i++ {
		universe, activeCubes = cycle(universe, activeCubes)
	}

	// Count active cubes
	return len(activeCubes)
}

func cycle(universe map[coord]rune, activeCubes []coord) (map[coord]rune, []coord) {
	// Create missing inactive cubes around active cubes
	for _, activeCube := range activeCubes {
		createMissingNeighbours(universe, activeCube)
	}

	// In a new universe, set updated states
	newUniverse := make(map[coord]rune)
	var newActiveCubes []coord

	for cube, value := range universe {
		activeNeighbours := countActiveNeighbours(universe, cube)
		if value == '#' && (activeNeighbours == 2 || activeNeighbours == 3) {
			newUniverse[cube] = '#'
			newActiveCubes = append(newActiveCubes, cube)
		} else if value == '.' && activeNeighbours == 3 {
			newUniverse[cube] = '#'
			newActiveCubes = append(newActiveCubes, cube)
		} else {
			newUniverse[cube] = '.'
		}
	}

	return newUniverse, newActiveCubes
}

func createMissingNeighbours(universe map[coord]rune, cube coord) {
	for _, neighbour := range aroundCoords(cube) {
		if _, exists := universe[neighbour]; !exists {
			universe[neighbour] = '.'
		}
	}
}

func countActiveNeighbours(universe map[coord]rune, cube coord) int {
	sumActive := 0
	for _, neighbour := range aroundCoords(cube) {
		if value, exists := universe[neighbour]; exists && value == '#' {
			sumActive++
		}
	}
	return sumActive
}

func aroundCoords(cube coord) []coord {
	var neighbours []coord
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				if !(dx == 0 && dy == 0 && dz == 0) {
					neighbours = append(neighbours, coord{cube.x+dx, cube.y+dy, cube.z+dz})
				}
			}
		}
	}
	return neighbours
}

/**
	PART 2
 */
func Part2() int {
	universe, activeCubes := common4d()
	cycleNumber := 6

	for i := 0; i < cycleNumber; i++ {
		universe, activeCubes = cycle4d(universe, activeCubes)
	}

	// Count active cubes
	return len(activeCubes)
}

func cycle4d(universe map[coord4d]rune, activeCubes []coord4d) (map[coord4d]rune, []coord4d) {
	// Create missing inactive cubes around active cubes
	for _, activeCube := range activeCubes {
		createMissingNeighbours4d(universe, activeCube)
	}

	// In a new universe, set updated states
	newUniverse := make(map[coord4d]rune)
	var newActiveCubes []coord4d

	for cube, value := range universe {
		activeNeighbours := countActiveNeighbours4d(universe, cube)
		if value == '#' && (activeNeighbours == 2 || activeNeighbours == 3) {
			newUniverse[cube] = '#'
			newActiveCubes = append(newActiveCubes, cube)
		} else if value == '.' && activeNeighbours == 3 {
			newUniverse[cube] = '#'
			newActiveCubes = append(newActiveCubes, cube)
		} else {
			newUniverse[cube] = '.'
		}
	}

	return newUniverse, newActiveCubes
}

func createMissingNeighbours4d(universe map[coord4d]rune, cube coord4d) {
	for _, neighbour := range aroundCoords4d(cube) {
		if _, exists := universe[neighbour]; !exists {
			universe[neighbour] = '.'
		}
	}
}

func countActiveNeighbours4d(universe map[coord4d]rune, cube coord4d) int {
	sumActive := 0
	for _, neighbour := range aroundCoords4d(cube) {
		if value, exists := universe[neighbour]; exists && value == '#' {
			sumActive++
		}
	}
	return sumActive
}

func aroundCoords4d(cube coord4d) []coord4d {
	var neighbours []coord4d
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			for dz := -1; dz <= 1; dz++ {
				for dw := -1; dw <= 1; dw++ {
					if !(dx == 0 && dy == 0 && dz == 0 && dw == 0) {
						neighbours = append(neighbours, coord4d{cube.x+dx, cube.y+dy, cube.z+dz, cube.w+dw})
					}
				}
			}
		}
	}
	return neighbours
}
package day12

import (
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type Ship struct {
	direction	int
	x			int
	y			int
}

type Waypoint struct {
	x			int
	y			int
}

type Command interface {
	apply(ship *Ship)
}

type North struct {
	units		int
}

type South struct {
	units		int
}

type East struct {
	units		int
}

type West struct {
	units		int
}

type Left struct {
	units		int
}

type Right struct {
	units		int
}

type Forward struct {
	units		int
}

func common() []Command {
	content, err := ioutil.ReadFile("2020/inputs/day12")
	if err != nil {
		log.Fatal(err)
	}

	var commands []Command
	for _, line := range strings.Split(string(content), "\n") {
		commands = append(commands, toCommand(line))
	}

	return commands
}

func toCommand(line string) Command {
	units, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(err)
	}

	switch string(line[0]) {
	case "N": return North { units }
	case "S": return South { units }
	case "E": return East { units }
	case "W": return West { units }
	case "L": return Left { units }
	case "R": return Right { units }
	case "F": return Forward { units }
	default: panic("unknown command: " + line)
	}
}

func Part1() int {
	ship := Ship{}
	commands := common()

	for _, command := range commands {
		command.apply(&ship)
	}

	return int(math.Abs(float64(ship.x)) + math.Abs(float64(ship.y)))
}

func Part2() int {
	ship := Ship{}
	waypoint := Waypoint { ship.x + 10, ship.y + 1 }
	commands := common()
}

/**
	ALL COMMANDS
 */

func (north North) apply(ship *Ship) {
	ship.y += north.units
}

func (south South) apply(ship *Ship) {
	ship.y -= south.units
}

func (east East) apply(ship *Ship) {
	ship.x += east.units
}

func (west West) apply(ship *Ship) {
	ship.x -= west.units
}

func (left Left) apply(ship *Ship) {
	ship.direction = (ship.direction + left.units) % 360
	if ship.direction < 0 {
		ship.direction += 360
	}
}

func (right Right) apply(ship *Ship) {
	ship.direction = (ship.direction - right.units) % 360
	if ship.direction < 0 {
		ship.direction += 360
	}
}

func (forward Forward) apply(ship *Ship) {
	switch ship.direction {
	case 0: 	ship.x += forward.units
	case 90: 	ship.y += forward.units
	case 180:	ship.x -= forward.units
	case 270:	ship.y -= forward.units
	default: panic("unknown direction: " + strconv.Itoa(ship.direction))
	}
}
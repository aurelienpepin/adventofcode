package day8

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction interface {
	execute(line int, acc int) (int, int)
}

type Jump struct {
	value int
}

type Acc struct {
	value int
}

type Nop struct {
	value int
}

func common() []string {
	content, err := ioutil.ReadFile("2020/inputs/day8")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}

func Part1() int {
	visited := make(map[int]bool)
	line, acc := 0, 0

	input := common()
	var program []Instruction

	for _, inputLine := range input {
		program = append(program, toInstruction(inputLine))
	}

	for {
		if _, alreadyVisited := visited[line]; alreadyVisited {
			return acc
		} else {
			visited[line] = true
			line, acc = program[line].execute(line, acc)
		}
	}
}

func toInstruction(input string) Instruction {
	parts := strings.Split(input, " ")
	value, err := strconv.Atoi(parts[1][1:])
	if err != nil {
		log.Fatal(err)
	}

	if string(parts[1][0]) == "-" {
		value = value * -1
	}

	switch parts[0] {
	case "jmp":
		return Jump{value}
	case "acc":
		return Acc{value}
	case "nop":
		return Nop{value}
	default:
		panic("unknown instruction")
	}
}

func (jump Jump) execute(line int, acc int) (int, int) {
	return line + jump.value, acc
}

func (ac Acc) execute(line int, acc int) (int, int) {
	return line + 1, acc + ac.value
}

func (nop Nop) execute(line int, acc int) (int, int) {
	return line + 1, acc
}
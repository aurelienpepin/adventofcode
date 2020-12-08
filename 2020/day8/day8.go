package day8

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Instruction interface {
	execute(line int, acc int) 	(int, int)
	getType()					string
	getValue()					int
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

func Part2() int {
	input := common()
	var program []Instruction

	for _, inputLine := range input {
		program = append(program, toInstruction(inputLine))
	}

	// Try to switch all instructions
	for i, instruction := range program {
		reversed, reversible := reverse(instruction)
		if reversible {
			program[i] = reversed
			acc, end := executeProgram(program)
			if end {
				return acc
			}

			program[i] = instruction
		}
	}

	panic("not found")
}

func executeProgram(program []Instruction) (acc int, end bool) {
	line, visited := 0, make(map[int]bool)

	for {
		if line == len(program) {
			return acc, true
		}

		if _, alreadyVisited := visited[line]; alreadyVisited {
			return acc, false
		} else {
			visited[line] = true
			line, acc = program[line].execute(line, acc)
		}
	}
}

func reverse(instruction Instruction) (Instruction, bool) {
	switch instruction.getType() {
	case "acc":
		return instruction, false
	case "jump":
		return Nop{instruction.getValue()}, true
	case "nop":
		return Jump{instruction.getValue()}, true
	default:
		panic("unknown instruction")
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

func (jump Jump) getType() string {
	return "jump"
}

func (jump Jump) getValue() int {
	return jump.value
}

func (ac Acc) execute(line int, acc int) (int, int) {
	return line + 1, acc + ac.value
}

func (ac Acc) getType() string {
	return "acc"
}

func (ac Acc) getValue() int {
	return ac.value
}

func (nop Nop) execute(line int, acc int) (int, int) {
	return line + 1, acc
}

func (nop Nop) getType() string {
	return "nop"
}

func (nop Nop) getValue() int {
	return nop.value
}
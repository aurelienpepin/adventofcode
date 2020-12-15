package day14

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Part1() uint64 {
	content, err := ioutil.ReadFile("2020/inputs/day14")
	if err != nil {
		log.Fatal(err)
	}

	memory := make(map[string]uint64)
	var currentMask string

	for _, line := range strings.Split(string(content), "\n") {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			currentMask = parts[1]
		} else {
			if value, err := strconv.Atoi(parts[1]); err == nil {
				memory[parts[0]] = forceBits(uint64(value), currentMask)
			} else {
				panic("bad value for memory")
			}
		}
	}

	// Sum the content of the memory
	sum := uint64(0)
	for _, value := range memory {
		sum += value
	}
	return sum
}

func forceBits(value uint64, mask string) uint64 {
	for i, bit := range mask {
		switch string(bit) {
		case "1":
			value = value | uint64(1 << (len(mask) - i - 1))
		case "0":
			value = value & uint64(^(1 << (len(mask) - i - 1)))
		default:
			// ignore
		}
		// fmt.Printf("%08b\n", value)
	}

	return value
}
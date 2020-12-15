package day14

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func common() []string {
	content, err := ioutil.ReadFile("2020/inputs/day14")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}

func Part1() uint64 {
	memory := make(map[string]uint64)
	var currentMask string

	for _, line := range common() {
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

func Part2() uint64 {
	memory := make(map[uint64]uint64)
	var currentMask string

	for _, line := range common() {
		parts := strings.Split(line, " = ")
		if parts[0] == "mask" {
			currentMask = parts[1]
		} else {
			left, errLeft := strconv.Atoi(parts[0][4:len(parts[0]) - 1])
			right, errRight := strconv.Atoi(parts[1])

			if errLeft == nil && errRight == nil {
				for _, address := range forceBinaryAddresses(uint64(left), currentMask) {
					memory[address] = uint64(right)
				}
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

func forceBinaryAddresses(value uint64, mask string) []uint64 {
	return combineAddresses(value, mask, len(mask) - 1)
}

func combineAddresses(value uint64, mask string, index int) []uint64 {
	invIndex := len(mask) - index - 1

	if index == 0 {
		var addresses []uint64
		switch mask[invIndex] {
		case '0':
			addresses = append(addresses, checkBit(value, 0))
		case '1':
			addresses = append(addresses, 1)
		case 'X':
			addresses = append(addresses, 0)
			addresses = append(addresses, 1)
		}

		return addresses
	}

	smallerAddresses := combineAddresses(value, mask, index - 1)
	switch mask[invIndex] {
	case '0':
		if checkBit(value, index) == 1 {
			for i, _ := range smallerAddresses {
				smallerAddresses[i] += 1 << index
			}
		}
	case '1':
		for i, _ := range smallerAddresses {
			smallerAddresses[i] += 1 << index
		}
	case 'X':
		n := len(smallerAddresses)
		for i := 0; i < n; i++ {
			smallerAddresses = append(smallerAddresses, smallerAddresses[i])
			smallerAddresses[i] += 1 << index
		}
	}

	return smallerAddresses
}

func checkBit(value uint64, index int) uint64 {
	return (value >> index) & uint64(1)
}
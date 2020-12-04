package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("2020/inputs/day4")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(splitAt("\n\n"))
	var passports []string

	for scanner.Scan() {
		passports = append(passports, scanner.Text())
	}

	fmt.Println(part1_day4(passports))
}

func part1_day4(passports []string) int {
	n := 0
	for _, p := range passports {
		if isValid(p) {
			n++
		}
	}
	return n
}

var allFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid", "cid"}
var mandatoryFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func isValid(passport string) bool {
	fields := strings.FieldsFunc(passport, splitter)
	index := make(map[string]string)
	set := make(map[string]bool)

	for _, field := range fields {
		parts := strings.Split(field, ":")
		index[parts[0]] = parts[1]
		set[parts[0]] = true
	}

	if len(fields) < len(mandatoryFields) || len(fields) > len(allFields) {
		return false
	}

	for _, field := range mandatoryFields {
		if value, ok := index[field]; !ok || !isConform(field, value) {
			return false
		}
	}

	if _, ok := set["cid"]; len(fields) == len(allFields) && !ok {
		return false
	}
	return true
}

func isConform(key string, value string) bool {
	switch key {
	case "byr":
		return isInInterval(value, 1920, 2002)
	case "iyr":
		return isInInterval(value, 2010, 2020)
	case "eyr":
		return isInInterval(value, 2020, 2030)
	case "hgt":
		return isHeight(value)
	case "hcl":
		return isColor(value)
	case "ecl":
		return stringInSlice(value, []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"})
	case "pid":
		_, err := strconv.Atoi(value)
		return len(value) == 9 && err == nil
	default:
		return true
	}
}

func isInInterval(value string, min int, max int) bool {
	v, err := strconv.Atoi(value)
	return err == nil && v >= min && v <= max
}

func isColor(value string) bool {
	if len(value) != 7 {
		return false
	}

	if string(value[0]) != "#" {
		return false
	}

	for i := 1; i < len(value); i++ {
		if !stringInSlice(string(value[i]), []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}) {
			return false
		}
	}

	return true
}

func isHeight(value string) bool {
	if len(value) < 3 {
		return false
	}

	unit := value[len(value) - 2:]
	if unit == "cm" {
		return isInInterval(value[:len(value) - 2], 150, 193)
	} else if unit == "in" {
		return isInInterval(value[:len(value) - 2], 59, 76)
	}
	return false
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func splitter(r rune) bool {
	return r == '\n' || r == ' '
}

// https://stackoverflow.com/a/57232670
func splitAt(substring string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte(substring)
	searchLen := len(searchBytes)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		// Return nothing if at end of file and no data passed
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		// Find next separator and return token
		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return dataLen, data, nil
		}

		// Request more data.
		return 0, nil, nil
	}
}
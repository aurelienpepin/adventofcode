package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
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
	//index := make(map[string]string)
	set := make(map[string]bool)

	for _, field := range fields {
		parts := strings.Split(field, ":")
		//index[parts[0]] = parts[1]
		set[parts[0]] = true
	}

	if len(fields) < len(mandatoryFields) || len(fields) > len(allFields) {
		return false
	}

	for _, field := range mandatoryFields {
		if _, ok := set[field]; !ok {
			return false
		}
	}

	if _, ok := set["cid"]; len(fields) == len(allFields) && !ok {
		return false
	}
	return true
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
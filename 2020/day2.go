package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Policy struct {
	key1     int
	key2     int
	target   string
	password string
}

func main() {
	file, err := os.Open("2020/inputs/day2")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var policies []Policy

	for scanner.Scan() {
		content := strings.Split(scanner.Text(), ": ")
		rules := strings.FieldsFunc(content[0], rulesSplitter)

		minRule, err1 := strconv.Atoi(rules[0])
		maxRule, err2 := strconv.Atoi(rules[1])
		if err1 != nil {
			log.Fatal(err1)
		}

		if err2 != nil {
			log.Fatal(err2)
		}

		policies = append(policies, Policy{
			key1:     minRule,
			key2:     maxRule,
			target:   rules[2],
			password: content[1],
		})
	}

	score := 0
	for _, p := range policies {
		// if p.isValid() {
		if p.isValid2() {
			score++
		}
	}

	fmt.Println(score)
}

func (policy Policy) isValid() bool {
	count := strings.Count(policy.password, policy.target)
	return policy.key1 <= count && count <= policy.key2
}

func (policy Policy) isValid2() bool {
	r1 := string(policy.password[policy.key1-1]) == policy.target
	r2 := string(policy.password[policy.key2-1]) == policy.target
	return r1 != r2 // xor
}

func rulesSplitter(r rune) bool {
	return r == ' ' || r == '-'
}

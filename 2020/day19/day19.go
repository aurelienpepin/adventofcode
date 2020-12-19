package day19

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type rule interface {
	isUnit()	bool
	getLeft()	int
	getTerm()	rune
	getRight1()	int
	getRight2()	int
}

type productionRule struct {
	left	int
	right1	int
	right2	int
}

type unitProductionRule struct {
	left	int
	right	rune
}

func common() ([]rule, []string) {
	content, err := ioutil.ReadFile("2020/inputs/day19")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(content), "\n\n")
	return commonRule(parts[0]), strings.Split(parts[1], "\n")
}

func commonRule(s string) []rule {
	lines := strings.Split(s, "\n")
	lines = decomposePipedRules(lines)
	lines = removeDirectRules(lines)

	var rules []rule
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rightsParts := strings.Split(parts[1], " ")

		if len(rightsParts) == 1 {
			rules = append(rules, unitProductionRule{
				left:  toInteger(parts[0]),
				right: toRune(parts[1]),
			})
		} else {
			rules = append(rules, productionRule{
				left:   toInteger(parts[0]),
				right1: toInteger(rightsParts[0]),
				right2: toInteger(rightsParts[1]),
			})
		}
	}

	//for _, rule := range rules {
	//	fmt.Printf("%+v\n", rule)
	//}
	return rules
}

/**
	PART 1
 */
func Part1() int {
	rules, sentences := common()
	numberOfValidSentences := 0
	ruleIndex := index(rules)

	for _, sentence := range sentences {
		if cyk(sentence, rules, ruleIndex) {
			numberOfValidSentences++
			fmt.Println("yes:", sentence)
		} else {
			fmt.Println("no:", sentence)
		}
	}

	return numberOfValidSentences
}

func index(rules []rule) map[int][]int {
	ruleIndex := make(map[int][]int)
	for i, rule := range rules {
		ruleIndex[rule.getLeft()] = append(ruleIndex[rule.getLeft()], i)
	}
	return ruleIndex
}

/**
	CYK ALGORITHM
 */
func cyk(sentence string, rules []rule, ruleIndex map[int][]int) bool {
	p, n := initP(sentence, rules), len(sentence)

	for s := 0; s < n; s++ {
		for v, rule := range rules {
			if rule.isUnit() && rune(sentence[s]) == rule.getTerm() {
				p[1][s][v] = true
			}
		}
	}

	for l := 2; l <= n; l++ {
		for s := 0; s <= n - l; s++ {
			for pa := 1; pa <= l - 1; pa++ {
				for a, rule := range rules {
					if !rule.isUnit() {

						// Now for all combinations B x C in rule A -> BC
						for _, B := range ruleIndex[rule.getRight1()] {
							for _, C := range ruleIndex[rule.getRight2()] {
								if p[pa][s][B] && p[l-pa][s+pa][C] {
									p[l][s][a] = true
								}
							}
						}
					}
				}
			}
		}
	}

	return p[n][0][ruleIndex[0][0]]
}

func initP(sentence string, rules []rule) [][][]bool {
	var p [][][]bool
	for i := 0; i <= len(sentence); i++ {

		var pp [][]bool
		for j := 0; j < len(sentence); j++ {
			ppp := make([]bool, len(rules))
			pp = append(pp, ppp)
		}
		p = append(p, pp)
	}

	return p
}

/**
	GRAMMAR SIMPLIFICATION
 */
func decomposePipedRules(lines []string) []string {
	var newLines []string
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rightParts := strings.Split(parts[1], " | ")

		if len(rightParts) == 1 { // Normal rule
			newLines = append(newLines, line)
		} else { // Piped rule
			for _, rightPart := range rightParts {
				newLines = append(newLines, parts[0] + ": " + rightPart)
			}
		}
	}
	return newLines
}

func removeDirectRules(lines []string) []string {
	var newLines, newNewLines []string
	replacements := make(map[string][]string)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		if rightParts := strings.Split(parts[1], " "); len(rightParts) == 1 && isInteger(rightParts[0]) {
			replacements[parts[0]] = append(replacements[parts[0]], parts[1])
		} else {
			newLines = append(newLines, line)
		}
	}

	// Apply replacements
	for _, line := range newLines {
		for _, replacedLine := range replaceAll(line, replacements) {
			newNewLines = append(newNewLines, replacedLine)
		}
	}

	return newNewLines
}

func replaceAll(line string, replacements map[string][]string) []string {
	parts := strings.Split(line, " ")
	updatedLines := []string{parts[0]}

	for i := 1; i < len(parts); i++ {
		if replacerList, ok := replacements[parts[i]]; ok {
			var changingLines []string

			for _, line := range updatedLines {
				for _, replacer := range replacerList {
					changingLines = append(changingLines, line + " " + replacer)
				}
			}

			updatedLines = changingLines
		} else {
			for j := 0; j < len(updatedLines); j++ {
				updatedLines[j] += " " + parts[i]
			}
		}
	}

	return updatedLines
}

func isDirectRule(line string) (int, bool) {
	parts := strings.Split(line, ": ")
	if val, err := strconv.Atoi(parts[1]); len(parts) == 2 && err == nil {
		return val, true
	}

	return -1, false
}

func isInteger(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func toInteger(input string) int {
	value, err := strconv.Atoi(input)
	if err != nil {
		panic("not an integer: " + input)
	}
	return value
}

func toRune(input string) rune {
	return rune(input[1])
}

func (p productionRule) isUnit() bool {
	return false
}

func (p productionRule) getLeft() int {
	return p.left
}

func (p productionRule) getTerm() rune {
	panic("no term!")
}

func (p productionRule) getRight1() int {
	return p.right1
}

func (p productionRule) getRight2() int {
	return p.right2
}

func (u unitProductionRule) isUnit() bool {
	return true
}

func (u unitProductionRule) getLeft() int {
	return u.left
}

func (u unitProductionRule) getTerm() rune {
	return u.right
}

func (u unitProductionRule) getRight1() int {
	panic("no right1!")
}

func (u unitProductionRule) getRight2() int {
	panic("no right2!")
}
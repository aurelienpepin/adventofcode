package day18

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Expression interface {
	evaluate()	int
}

type IntExpression struct {
	value		int
}

type PlusExpression struct {
	left		Expression
	right		Expression
}

type MultExpression	struct {
	left		Expression
	right		Expression
}

type ParenthExpression struct {
	expr		Expression
}

func common() []Expression {
	content, err := ioutil.ReadFile("2020/inputs/day18")
	if err != nil {
		log.Fatal(err)
	}

	var expressions []Expression
	for _, line := range strings.Split(string(content), "\n") {
		expressions = append(expressions, parse(line))
	}

	return expressions
}

func Part1() int {
	sum := 0
	expressions := common()

	for _, expression := range expressions {
		sum += expression.evaluate()
	}

	return sum
}

func parse(formula string) Expression {
	if len(formula) == 1 {
		value, _ := strconv.Atoi(string(formula[0]))
		return IntExpression{value}
	}

	if isDigit(last(formula)) {
		if beforeLast(formula) == '+' {
			return PlusExpression{parse(formula[:len(formula) - 4]), parse(string(last(formula)))}
		} else if beforeLast(formula) == '*' {
			return MultExpression{parse(formula[:len(formula) - 4]), parse(string(last(formula)))}
		} else {
			panic("unknown beforeLast: " + string(beforeLast(formula)))
		}
	} else { // Last is )
		open := matchParenthesis(formula)
		if open == 0 {
			return ParenthExpression{parse(formula[1:len(formula) - 1])}
		} else {
			op := formula[open - 2]
			if op == '+' {
				return PlusExpression{parse(formula[:open - 3]), ParenthExpression{parse(formula[open+1:len(formula) - 1])}}
			} else if op == '*' {
				return MultExpression{parse(formula[:open - 3]), ParenthExpression{parse(formula[open+1:len(formula) - 1])}}
			} else {
				panic("unknown: " + string(op))
			}
		}
	}
}

func matchParenthesis(formula string) int {
	counter := 0
	for i := len(formula) - 1; i >= 0; i-- {
		if formula[i] == ')' {
			counter--
		} else if formula[i] == '(' {
			counter++
		}

		if counter == 0 {
			return i
		}
	}

	panic("no matching parenthesis in:" + formula)
}

func last(formula string) rune {
	return rune(formula[len(formula) - 1])
}

func beforeLast(formula string) rune {
	return rune(formula[len(formula) - 3])
}

func isDigit(char rune) bool {
	return char >= '1' && char <= '9'
}

func (i IntExpression) evaluate() int {
	return i.value
}

func (p PlusExpression) evaluate() int {
	valLeft := p.left.evaluate()
	valRight := p.right.evaluate()
	return valLeft + valRight
}

func (m MultExpression) evaluate() int {
	valLeft := m.left.evaluate()
	valRight := m.right.evaluate()
	return valLeft * valRight
}

func (p ParenthExpression) evaluate() int {
	return p.expr.evaluate()
}


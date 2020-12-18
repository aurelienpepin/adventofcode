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

func common(parsingFunction func(string) Expression) []Expression {
	content, err := ioutil.ReadFile("2020/inputs/day18")
	if err != nil {
		log.Fatal(err)
	}

	var expressions []Expression
	for _, line := range strings.Split(string(content), "\n") {
		expressions = append(expressions, parsingFunction(line))
	}

	return expressions
}

func Part1() int {
	sum := 0
	expressions := common(parse)

	for _, expression := range expressions {
		sum += expression.evaluate()
	}

	return sum
}

func Part2() int64 {
	sum := int64(0)
	expressions := common(parse2)

	for _, expression := range expressions {
		sum += int64(expression.evaluate())
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

func parse2(formula string) Expression {
	// Modify the input formula by surrounding all + with ()
	i := 0

	for i != len(formula) {
		if formula[i] == '+' {
			formula = surround(formula, i)
			i++
		}

		i++
	}

	return parse(formula)
}

func surround(formula string, plusIndex int) string {
	// Left part
	if isDigit(rune(formula[plusIndex-2])) {
		formula = formula[:plusIndex-2] + "(" + formula[plusIndex-2:]
	} else {
		open := matchParenthesis(formula[:plusIndex-1])
		formula = formula[:open] + "(" + formula[open:]
	}

	// Right part
	plusIndex++
	if isDigit(rune(formula[plusIndex+2])) {
		formula = formula[:plusIndex+3] + ")" + formula[plusIndex+3:]
	} else {
		closed := matchParenthesisRight(formula[plusIndex + 2:])
		// fmt.Println(closed, "/", formula, "/", formula[plusIndex + 2:], "/")
		formula = formula[:plusIndex + 2+closed] + ")" + formula[plusIndex + 2+closed:]
	}
	return formula
}

func matchParenthesisRight(formula string) int {
	counter := 0
	for i := 0; i < len(formula); i++ {
		if formula[i] == ')' {
			counter--
		} else if formula[i] == '(' {
			counter++
		}

		if counter == 0 {
			return i
		}
	}

	panic("no matching parenthesis right in:" + formula)
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


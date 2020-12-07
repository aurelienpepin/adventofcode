package day7

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Graph struct {
	adj		map[string][]string
}

type Bag struct {
	name		string
	quantity	int
	innerBags	[]Bag
}

func (g *Graph) Add(a string, b string) {
	g.adj[a] = append(g.adj[a], b)
}

func (g Graph) AllNodesReachableFrom(source string) map[string]bool {
	visited := make(map[string]bool)
	var stack []string
	stack = append(stack, source)

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		for _, v := range g.adj[node] {
			if _, ok := visited[v]; !ok {
				visited[v] = true
				stack = append(stack, v)
			}
		}
	}

	return visited
}

func common() []string {
	content, err := ioutil.ReadFile("2020/inputs/day7")
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(content), "\n")
}

func Part1() int {
	g := Graph{ make(map[string][]string) }
	inputLines := common()

	for _, line := range inputLines {
		bag := extractBag(line)
		for _, innerBag := range bag.innerBags {
			g.Add(innerBag.name, bag.name)
		}
	}

	fmt.Println(g)
	return len(g.AllNodesReachableFrom("shiny gold"))
}

func extractBag(input string) Bag {
	parts := strings.Split(input, " ")
	bagName := parts[0] + " " + parts[1]
	var innerBags []Bag

	for i := 4; i < len(parts); i += 4 {
		quantity, err := strconv.Atoi(parts[i])
		if err != nil {
			continue // quantity is: "no"
		}

		innerBags = append(innerBags, Bag{
			name:      parts[i + 1] + " " + parts[i + 2],
			quantity:  quantity,
			innerBags: nil,
		})
	}

	return Bag { bagName, 1, innerBags }
}
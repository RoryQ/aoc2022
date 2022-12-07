package main

import (
	"bufio"
	_ "embed"
	"sort"
	"strconv"
	"strings"

	"github.com/kr/pretty"
	"github.com/samber/lo"
)

//go:embed input
var input string

func main() {
	one()
	two()
}

func two() {
	root := parseInput()

	total := 70000000 - root.Total()
	target := 30000000 - total
	atLeast := findAtLeast(&root, target)
	sort.SliceStable(atLeast, func(i, j int) bool {
		return atLeast[i].Total() < atLeast[j].Total()
	})

	pretty.Println(atLeast[0].Total())
}

type Node struct {
	Name     string
	IsDir    bool
	Size     int
	Children map[string]*Node
}

func one() {
	root := parseInput()

	smallest := findAtMost(&root, 100000)
	println(lo.SumBy(smallest, func(item *Node) int { return item.Total() }))
}

func parseInput() Node {
	scan := bufio.NewScanner(strings.NewReader(input))

	scan.Scan() // pop first

	name := strings.Split(scan.Text(), " ")[2]
	root := Node{Name: name, Children: map[string]*Node{}, IsDir: true}
	current := &root
	stack := []*Node{}
	for scan.Scan() {
	Next:
		line := scan.Text()
		if isCommand(line, "cd ..") {
			current, stack = popStack(stack)
		} else if isCommand(line, "cd") {
			name := strings.Split(line, " ")[2]
			stack = append(stack, current)
			current = current.Children[name]
		} else if isCommand(line, "ls") {
			for scan.Scan() {
				line := scan.Text()
				if isCommand(line, "") {
					goto Next
				}
				split := strings.Split(line, " ")
				size, err := strconv.Atoi(split[0])
				name := split[1]
				current.Children[name] = &Node{Name: name, Size: size, IsDir: err != nil, Children: map[string]*Node{}}
			}
		}
	}
	return root
}

func (n *Node) Total() int {
	return n.Size + lo.SumBy(lo.Values(n.Children), func(item *Node) int { return item.Total() })
}

func findAtLeast(root *Node, size int) []*Node {
	found := []*Node{}
	toVisit := []*Node{root}
	for len(toVisit) > 0 {
		var node *Node
		node, toVisit = popStack(toVisit)
		if node.IsDir {
			toVisit = append(toVisit, lo.Values(node.Children)...)

			if node.Total() >= size {
				found = append(found, node)
			}
		}
	}

	return found
}

func findAtMost(root *Node, size int) []*Node {
	found := []*Node{}
	toVisit := []*Node{root}
	for len(toVisit) > 0 {
		var node *Node
		node, toVisit = popStack(toVisit)
		if node.IsDir {
			toVisit = append(toVisit, lo.Values(node.Children)...)

			if node.Total() <= size {
				found = append(found, node)
			}
		}
	}

	return found
}

func popStack(stack []*Node) (*Node, []*Node) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func isCommand(line, s string) bool {
	return strings.HasPrefix(line, "$") && strings.Contains(line, s)
}

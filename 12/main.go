package main

import (
	"bufio"
	_ "embed"
	"math"
	"strings"

	"github.com/kr/pretty"
	"github.com/samber/lo"
)

var (
	//go:embed input
	input string
	//go:embed input_example
	test string

	println = pretty.Println
)

func main() {
	one(bufio.NewScanner(strings.NewReader(test)))
	one(bufio.NewScanner(strings.NewReader(input)))

	two(bufio.NewScanner(strings.NewReader(test)))
	two(bufio.NewScanner(strings.NewReader(input)))
}

func one(scanner *bufio.Scanner) {
	m := parseInput(scanner)
	m.findPath([]tup{{coord: m.start}})
	println(m.grid[m.end.A][m.end.B].distance)
}

func two(scanner *bufio.Scanner) {
	m := parseInput(scanner)
	queue := []tup{}
	for _, items := range m.grid {
		for _, item := range items {
			if item.height == 0 {
				queue = append(queue, tup{coord: item.coord})
			}
		}
	}
	m.findPath(queue)
	println(m.grid[m.end.A][m.end.B].distance)
}

type coord = lo.Tuple2[int, int]

type Item struct {
	distance int
	height   int
	coord
}

type Map struct {
	grid  [][]Item
	start coord
	end   coord
}

type tup struct {
	coord    coord
	distance int
}

func parseInput(scanner *bufio.Scanner) Map {
	m := Map{grid: [][]Item{}}
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]Item, 0, len(line))
		for i, c := range line {
			coordinate := coord{A: len(m.grid), B: i}
			if c == 'E' {
				m.end = coordinate
			}
			if c == 'S' {
				m.start = coordinate
			}

			row = append(row, Item{
				coord:    coordinate,
				distance: math.MaxInt,
				height:   height(c),
			})
		}
		m.grid = append(m.grid, row)
	}
	return m
}

func (m *Map) findPath(queue []tup) {
	var cur tup
	for len(queue) > 0 {
		cur, queue = popFront(queue)

		if m.grid[cur.coord.A][cur.coord.B].distance < cur.distance {
			continue
		}

		if m.end == cur.coord {
			return
		}

		for _, next := range m.validDirections(cur.coord) {
			if m.grid[next.A][next.B].distance > cur.distance+1 {
				m.grid[next.A][next.B].distance = cur.distance + 1
				queue = append(queue, tup{next, cur.distance + 1})
			}
		}
	}
}

func (m Map) validDirections(p coord) []coord {
	var valid []coord
	getHeight := func(a, b int) int {
		return m.grid[a][b].height
	}
	cur := getHeight(p.A, p.B)

	check := func(dir coord) bool {
		return cur >= getHeight(dir.A, dir.B) || cur+1 == getHeight(dir.A, dir.B)
	}

	up := coord{p.A - 1, p.B}
	down := coord{p.A + 1, p.B}
	left := coord{p.A, p.B - 1}
	right := coord{p.A, p.B + 1}

	if up.A >= 0 && check(up) {
		valid = append(valid, up)
	}
	if down.A < len(m.grid) && check(down) {
		valid = append(valid, down)
	}
	if left.B >= 0 && check(left) {
		valid = append(valid, left)
	}
	if right.B < len(m.grid[0]) && check(right) {
		valid = append(valid, right)
	}

	return valid
}

func height(s rune) int {
	if s == 'S' {
		s = 'a'
	}
	if s == 'E' {
		s = 'z'
	}
	return int(s - 'a')
}

func popFront[T any](s []T) (T, []T) {
	return s[0], s[1:]
}

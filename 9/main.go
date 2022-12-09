package main

import (
	"bufio"
	_ "embed"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

//go:embed input
var input string

func main() {
	one()
	two()
}

func one() {
	scan := bufio.NewScanner(strings.NewReader(input))

	visited := []coord{{}}
	head, tail := coord{}, coord{}
	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")
		direction, count := line[0], lo.Must(strconv.Atoi(line[1]))

		switch direction {
		case "U":
			for i := 0; i < count; i++ {
				head.Y++
				if shouldUpdate(head, tail) {
					if diagonal(head, tail) {
						tail.Y = head.Y - 1
						tail.X = head.X
					} else {
						tail.Y++
					}
				}

				visited = append(visited, tail)
			}

		case "D":
			for i := 0; i < count; i++ {
				head.Y--
				if shouldUpdate(head, tail) {
					if diagonal(head, tail) {
						tail.Y = head.Y + 1
						tail.X = head.X
					} else {
						tail.Y--
					}
				}
				visited = append(visited, tail)
			}

		case "L":
			for i := 0; i < count; i++ {
				head.X--
				if shouldUpdate(head, tail) {
					if diagonal(head, tail) {
						tail.Y = head.Y
						tail.X = head.X + 1
					} else {
						tail.X--
					}
				}
				visited = append(visited, tail)
			}

		case "R":
			for i := 0; i < count; i++ {
				head.X++
				if shouldUpdate(head, tail) {
					if diagonal(head, tail) {
						tail.Y = head.Y
						tail.X = head.X - 1
					} else {
						tail.X++
					}
				}
				visited = append(visited, tail)
			}
		}
	}

	println(len(lo.Uniq(visited)))
}

type coord struct {
	X, Y int
}

func (c coord) Clone() *coord {
	return &c
}

func (rope Rope) UpdateBody() {
	for i := 1; i < len(rope); i++ {
		prev, next := rope[i-1], rope[i]
		updateNext(prev, next)
	}
}

type Rope []*coord

func two() {
	scan := bufio.NewScanner(strings.NewReader(input))

	visited := []coord{{}}
	rope := Rope(lo.Repeat(10, &coord{}))
	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")
		direction, count := line[0], lo.Must(strconv.Atoi(line[1]))

		switch direction {
		case "U":
			for i := 0; i < count; i++ {
				head, tail := rope[0], lo.Must(lo.Last(rope))
				head.Y++
				rope.UpdateBody()
				visited = append(visited, *tail)
			}

		case "D":
			for i := 0; i < count; i++ {
				head, tail := rope[0], lo.Must(lo.Last(rope))
				head.Y--
				rope.UpdateBody()
				visited = append(visited, *tail)
			}

		case "L":
			for i := 0; i < count; i++ {
				head, tail := rope[0], lo.Must(lo.Last(rope))
				head.X--
				rope.UpdateBody()
				visited = append(visited, *tail)
			}

		case "R":
			for i := 0; i < count; i++ {
				head, tail := rope[0], lo.Must(lo.Last(rope))
				head.X++
				rope.UpdateBody()
				visited = append(visited, *tail)
			}
		}
	}

	println(len(lo.Uniq(visited)))
}

func updateNext(head, tail *coord) {
	if shouldUpdate(*head, *tail) {
		tail.X += clamp(head.X - tail.X)
		tail.Y += clamp(head.Y - tail.Y)
	}
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func shouldUpdate(head, tail coord) bool {
	return abs(tail.Y-head.Y) > 1 || abs(tail.X-head.X) > 1

}
func diagonal(head, tail coord) bool {
	return tail.X != head.X && tail.Y != head.Y
}

func clamp(x int) int {
	if x > 0 {
		return 1
	}

	if x < 0 {
		return -1
	}
	return x
}

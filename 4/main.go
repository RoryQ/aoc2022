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

func two() {
	scan := bufio.NewScanner(strings.NewReader(input))
	total := 0
	for scan.Scan() {
		s := strings.Split(scan.Text(), ",")
		first, second := toRange(s[0]), toRange(s[1])

		intersect := lo.Intersect(first, second)
		if len(intersect) > 0 {
			total += 1
		}
	}
	println(total)
}

func one() {
	scan := bufio.NewScanner(strings.NewReader(input))
	total := 0
	for scan.Scan() {
		s := strings.Split(scan.Text(), ",")
		first, second := toRange(s[0]), toRange(s[1])

		intersect := lo.Intersect(first, second)
		if len(intersect) == len(first) || len(intersect) == len(second) {
			total += 1
		}
	}
	println(total)
}

func toRange(elf string) []int {
	s := strings.Split(elf, "-")
	start, end := lo.Must(strconv.Atoi(s[0])), lo.Must(strconv.Atoi(s[1]))
	return lo.RangeWithSteps(start, end+1, 1)
}

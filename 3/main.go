package main

import (
	"bufio"
	_ "embed"
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
	group := []string{}
	for scan.Scan() {
		s := scan.Text()
		group = append(group, s)
		if len(group) == 3 {
			c1, c2 := common(group[0], group[1]), common(group[1], group[2])
			c := common(c1, c2)
			total += valueof(c)
			group = []string{}
		}
	}
	println(total)
}

func one() {
	scan := bufio.NewScanner(strings.NewReader(input))
	total := 0
	for scan.Scan() {
		s := scan.Text()
		intersect := common(toCompartments(s))
		//println(intersect, valueof(intersect))
		total += valueof(intersect)
	}
	println(total)
}

func valueof(s string) int {
	v := int(rune(s[0]))

	if v > int('a') {
		return v - int('a') + 1
	}

	return v - int('A') + 27
}

func common(left, right string) string {
	i := lo.Intersect(strings.Split(left, ""), strings.Split(right, ""))
	return strings.Join(i, "")
}

func toCompartments(s string) (string, string) {
	return s[:len(s)/2], s[len(s)/2:]
}

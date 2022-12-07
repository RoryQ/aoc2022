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

type elf struct {
	number int
	food   []int
}

func main() {
	scanner := bufio.NewScanner(strings.NewReader(input))
	elves := []elf{}
	nextElf := elf{number: 1}
	for scanner.Scan() {
		if scanner.Text() == "" {
			elves = append(elves, nextElf)
			nextElf = elf{number: nextElf.number + 1}
			continue
		}
		food, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		nextElf.food = append(nextElf.food, food)
	}

	sort.Slice(elves, func(a int, b int) bool {
		return lo.Sum(elves[a].food) > lo.Sum(elves[b].food)
	})
	pretty.Println(elves[0], lo.Sum(elves[0].food))
	pretty.Println(elves[1], lo.Sum(elves[1].food))
	pretty.Println(elves[2], lo.Sum(elves[2].food))
	println(lo.Sum(elves[0].food) + lo.Sum(elves[1].food) + lo.Sum(elves[2].food))
}

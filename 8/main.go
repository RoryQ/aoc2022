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
	trees := grid()

	//pretty.Println(calcScore(trees, 1, 2))
	//pretty.Println(calcScore(trees, 0, 2))

	highest := 0
	for i := 0; i < len(trees); i++ {
		for j := 0; j < len(trees); j++ {
			highest = lo.Max([]int{calcScore(trees, i, j), highest})
		}
	}
	println(highest)
}

func calcScore(trees [][]int, a, b int) int {
	tree := trees[a][b]
	if a == 0 || b == 0 || a == len(trees)-1 || b == len(trees)-1 {
		return 0
	}

	down := 0
	for i, j := a+1, b; i < len(trees); i++ {
		down++
		if tree <= trees[i][j] {
			break
		}
	}
	right := 0
	for i, j := a, b+1; j < len(trees); j++ {
		right++
		if tree <= trees[i][j] {
			break
		}
	}
	left := 0
	for i, j := a, b-1; j >= 0; j-- {
		left++
		if tree <= trees[i][j] {
			break
		}
	}
	up := 0
	for i, j := a-1, b; i >= 0; i-- {
		up++
		if tree <= trees[i][j] {
			break
		}
	}

	return up * down * left * right
}

func one() {
	trees := grid()

	debug := [][]string{}
	for i := 0; i < len(trees); i++ {
		debug = append(debug, strings.Split(strings.Repeat(" ", len(trees)), ""))
	}

	visible := []coord{}
	for _, seq := range lineOfSightRanges(len(trees), len(trees)) {
		first := true
		var highest int
		for _, tup := range seq {
			tree := trees[tup.A][tup.B]

			if first {
				first = false
				debug[tup.A][tup.B] = "x"
				visible = append(visible, tup)
				highest = tree
				continue
			}

			if tree <= highest {
				continue
			}

			debug[tup.A][tup.B] = "x"
			visible = append(visible, tup)
			highest = lo.Max([]int{tree, highest})
		}
	}

	//for _, d := range debug { println(strings.Join(d, "")) }
	println(len(lo.Uniq(visible)))
}

func grid() [][]int {
	scan := bufio.NewScanner(strings.NewReader(input))

	trees := [][]int{}
	for scan.Scan() {
		line := scan.Text()
		trees = append(trees, lo.Map(strings.Split(line, ""), toInt))
	}
	return trees
}

type coord = lo.Tuple2[int, int]

func lineOfSightRanges(ilen, jlen int) [][]coord {
	//[][]int{
	//	{00, 01, 02, 03, 04},
	//	{10, 11, 12, 13, 14},
	//	{20, 21, 22, 23, 24},
	//	{30, 31, 32, 33, 34},
	//	{40, 41, 42, 43, 44},
	//} ij

	ranges := [][]coord{}

	for i := 0; i < ilen; i++ {
		ranges = append(ranges, lo.Zip2(repeat(i, jlen), lo.Range(jlen)))
	}

	for i := 0; i < ilen; i++ {
		ranges = append(ranges, lo.Zip2(repeat(i, jlen), lo.Reverse(lo.Range(jlen))))
	}

	for j := 0; j < jlen; j++ {
		ranges = append(ranges, lo.Zip2(lo.Range(ilen), repeat(j, jlen)))
	}

	for j := 0; j < jlen; j++ {
		ranges = append(ranges, lo.Zip2(lo.Reverse(lo.Range(ilen)), repeat(j, jlen)))
	}

	return ranges
}

func toInt(item string, _ int) int {
	return lo.Must(strconv.Atoi(item))
}

func repeat(num, times int) []int {
	s := make([]int, times)
	for i := 0; i < times; i++ {
		s[i] = num
	}
	return s
}

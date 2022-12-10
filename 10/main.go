package main

import (
	"bufio"
	_ "embed"
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

func one() {
	scan := bufio.NewScanner(strings.NewReader(input))

	stops := []int{20, 60, 100, 140, 180, 220}
	cycle := 0
	registerX := 1
	strengths := []int{}
	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")

		if line[0] == "noop" {
			cycle++
			if lo.Contains(stops, cycle) {
				strengths = append(strengths, registerX)
			}
			continue
		}

		cycle++
		if lo.Contains(stops, cycle) {
			strengths = append(strengths, registerX)
		}
		cycle++
		if lo.Contains(stops, cycle) {
			strengths = append(strengths, registerX)
		}
		registerX += atoi(line[1])
	}

	zipped := lo.Map(lo.Zip2(strengths, stops), func(item lo.Tuple2[int, int], _ int) int {
		return item.A * item.B
	})
	pretty.Println(lo.Sum(zipped))
}

func two() {
	scan := bufio.NewScanner(strings.NewReader(input))

	pixel := strings.Split(strings.Repeat(" ", 240), "")
	cycle := 1
	registerX := 1
	for scan.Scan() {
		line := strings.Split(scan.Text(), " ")

		if line[0] == "noop" {
			if abs(registerX-cycle%40) < 2 {
				pixel[cycle] = "#"
			}
			cycle++
			continue
		}

		if abs(registerX-cycle%40) < 2 {
			pixel[cycle] = "#"
		}
		cycle++
		registerX += atoi(line[1])
		if abs(registerX-cycle%40) < 2 {
			pixel[cycle] = "#"
		}
		cycle++
	}

	for i := 0; i < 240; i += 40 {
		pretty.Println(strings.Join(pixel[i:i+40], ""))
	}
}

func atoi(s string) int {
	return lo.Must(strconv.Atoi(s))
}

func abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

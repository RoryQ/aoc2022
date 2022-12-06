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

	for scan.Scan() {
		line := scan.Text()

		const target = 14
		for i := 0; i < len(line)-target; i++ {
			if len(lo.Uniq(strings.Split(line[i:i+target], ""))) == target {
				println(i + target)
				break
			}
		}
	}

}

func one() {
	scan := bufio.NewScanner(strings.NewReader(input))

	for scan.Scan() {
		line := scan.Text()

		for i := 0; i < len(line)-4; i++ {
			if len(lo.Uniq(strings.Split(line[i:i+4], ""))) == 4 {
				println(i + 4)
				break
			}
		}
	}

}

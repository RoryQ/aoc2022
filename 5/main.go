package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	. "github.com/kr/pretty"
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

	var stacks map[string][]string
	lines := []string{}
	for scan.Scan() {
		if scan.Text() == "" {
			stacks = parseInit(lines)
			break
		}
		lines = append(lines, scan.Text())
	}

	moves := []Move{}
	for scan.Scan() {
		text := scan.Text()
		moves = append(moves, parseMove(text))
	}

	for _, move := range moves {
		var took []string
		from := stacks[move.From]
		to := stacks[move.To]
		took, from = take(from, move.Count)
		stacks[move.From] = from
		stacks[move.To] = append(to, took...)
	}

	top := []string{}
	for i := 1; i <= len(stacks); i++ {
		top = append(top, lo.Must(lo.Last(stacks[strconv.Itoa(i)])))
	}
	Println(strings.Join(top, ""))
}

func one() {
	Println()
	scan := bufio.NewScanner(strings.NewReader(input))

	var stacks map[string][]string
	lines := []string{}
	for scan.Scan() {
		if scan.Text() == "" {
			stacks = parseInit(lines)
			break
		}
		lines = append(lines, scan.Text())
	}

	moves := []Move{}
	for scan.Scan() {
		text := scan.Text()
		moves = append(moves, parseMove(text))
	}

	for _, move := range moves {
		var popped []string
		from := stacks[move.From]
		to := stacks[move.To]
		popped, from = pop(from, move.Count)
		stacks[move.From] = from
		stacks[move.To] = append(to, popped...)
	}

	top := []string{}
	for i := 1; i <= len(stacks); i++ {
		top = append(top, lo.Must(lo.Last(stacks[strconv.Itoa(i)])))
	}
	Println(strings.Join(top, ""))
}

func take(s []string, count int) ([]string, []string) {
	popped, rest := pop(s, count)
	lo.Reverse(popped)
	return popped, rest
}

func pop(s []string, count int) ([]string, []string) {
	popped := []string{}
	for i := 0; i < count; i++ {
		var last string
		last, s = s[len(s)-1], s[:len(s)-1]
		popped = append(popped, last)
	}

	return popped, s
}

type Move struct {
	Count    int
	From, To string
}

func parseMove(line string) Move {
	re := regexp.MustCompile(`move (?P<count>\d+) from (?P<from>\d+?) to (?P<to>\d+?)`)
	match := re.FindStringSubmatch(line)

	return Move{
		Count: lo.Must(strconv.Atoi(match[re.SubexpIndex("count")])),
		From:  match[re.SubexpIndex("from")],
		To:    match[re.SubexpIndex("to")],
	}
}

func parseInit(lines []string) map[string][]string {
	lines = lo.Reverse(lines)
	stacks := lo.SliceToMap(lo.WithoutEmpty(strings.Split(lines[0], " ")), func(item string) (string, []string) {
		return strings.TrimSpace(item), []string{}
	})

	padRight := func(str string, length int) string {
		return fmt.Sprintf("%-"+strconv.Itoa(length)+"s", str)
	}

	for _, line := range lines[1:] {
		line := padRight(line, len(lines[0]))
		for i := 1; i <= len(stacks); i++ {
			key := strconv.Itoa(i)
			s := stacks[key]
			spaces := i - 1
			squares := 3 * i
			middle := spaces + squares - 2
			letter := string(line[middle])
			if letter == " " {
				continue
			}
			stacks[key] = append(s, letter)
		}
	}

	return stacks
}

package main

import (
	"bufio"
	_ "embed"
	"sort"
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

var split = strings.Split

func one() {
	monkeys := getMonkeys()

	var item int
	for rounds := 0; rounds < 20; rounds++ {
		for _, monkey := range monkeys {
			for range monkey.Items {
				item, monkey.Items = popFront(monkey.Items)
				monkey.Inspected++
				worry := monkey.Operation(item) / 3
				if (worry % monkey.Test) == 0 {
					monkeys[monkey.TestTrue].AppendItem(worry)
				} else {
					monkeys[monkey.TestFalse].AppendItem(worry)
				}
			}
		}
	}

	sort.Slice(monkeys, func(a, b int) bool {
		return monkeys[a].Inspected > monkeys[b].Inspected
	})
	println(monkeys[0].Inspected * monkeys[1].Inspected)
}

type Monkey struct {
	Number    int
	Items     []int
	Operation func(old int) int
	Test      int
	TestTrue  int
	TestFalse int
	Inspected int
}

func (m *Monkey) AppendItem(i int) {
	m.Items = append(m.Items, i)
}

func toOperation(line []string) func(old int) int {
	if strings.Join(line, " ") == "old * old" {
		return func(old int) int {
			return old * old
		}
	}

	val := atoi(line[2])
	switch line[1] {
	case "*":
		return func(old int) int {
			return old * val
		}

	case "+":
		return func(old int) int {
			return old + val
		}
	}

	return nil
}

func atoi(s string) int {
	return lo.Must(strconv.Atoi(s))
}

func toInt(s string, _ int) int {
	return atoi(s)
}

func mapTrim(s string, _ int) string {
	return strings.Trim(s, ", ")
}

func popFront[T any](s []T) (T, []T) {
	return s[0], s[1:]
}

func lcm(x, y int) int {
	tmp := x
	for (tmp % y) != 0 {
		tmp += x
	}
	return tmp
}

func two() {
	monkeys := getMonkeys()

	testcap := lo.Reduce(monkeys, func(agg int, item *Monkey, _ int) int { return lcm(agg, item.Test) }, 1)
	var item int
	for rounds := 1; rounds <= 10_000; rounds++ {
		for _, monkey := range monkeys {
			for range monkey.Items {
				item, monkey.Items = popFront(monkey.Items)
				monkey.Inspected++
				worry := monkey.Operation(item) % testcap
				if (worry % monkey.Test) == 0 {
					monkeys[monkey.TestTrue].AppendItem(worry)
				} else {
					monkeys[monkey.TestFalse].AppendItem(worry)
				}
			}
		}
	}

	sort.Slice(monkeys, func(a, b int) bool {
		return monkeys[a].Inspected > monkeys[b].Inspected
	})

	println(monkeys[0].Inspected * monkeys[1].Inspected)
}

func getMonkeys() []*Monkey {
	scan := bufio.NewScanner(strings.NewReader(input))

	monkeys := []*Monkey{}
	for scan.Scan() {
		line := split(scan.Text(), " ")
		m := Monkey{
			Number: atoi(strings.TrimSuffix(line[1], ":")),
		}

		scan.Scan()
		line = split(split(scan.Text(), ":")[1], " ")
		m.Items = lo.Map(lo.WithoutEmpty(lo.Map(line, mapTrim)), toInt)

		scan.Scan()
		line = lo.WithoutEmpty(split(split(split(scan.Text(), ":")[1], "=")[1], " "))
		m.Operation = toOperation(line)

		scan.Scan()
		line = strings.Fields(scan.Text())
		m.Test = atoi(lo.Must(lo.Last(line)))

		scan.Scan()
		line = strings.Fields(scan.Text())
		m.TestTrue = atoi(lo.Must(lo.Last(line)))

		scan.Scan()
		line = strings.Fields(scan.Text())
		m.TestFalse = atoi(lo.Must(lo.Last(line)))
		scan.Scan()

		monkeys = append(monkeys, &m)
	}
	return monkeys
}

package main

import (
	"bufio"
	_ "embed"
	"strings"
)

//go:embed input
var input string

const (
	rock     = "rock"
	paper    = "paper"
	scissors = "scissors"
)

var strategy = map[string]string{
	"X": rock,
	"Y": paper,
	"Z": scissors,
}

func strategy2(opp, you string) string {
	if you == "X" {
		return loses(opp)
	}

	if you == "Y" {
		return draws(opp)
	}

	return beats(opp)
}

func beats(s string) string {
	if s == paper {
		return scissors
	}
	if s == rock {
		return paper
	}
	return rock
}

func draws(s string) string {
	return s
}

func loses(s string) string {
	if s == paper {
		return rock
	}
	if s == rock {
		return scissors
	}
	return paper
}

var opponent = map[string]string{
	"A": rock,
	"B": paper,
	"C": scissors,
}

func calcScore(opp, you string) int {
	score := 0
	if you == rock {
		score = 1
	}
	if you == paper {
		score = 2
	}
	if you == scissors {
		score = 3
	}

	return score + winner(opp, you)
}

func winner(opp, you string) int {
	switch opp {
	case rock:
		if you == paper {
			return 6
		}
		if you == scissors {
			return 0
		}
	case paper:
		if you == scissors {
			return 6
		}
		if you == rock {
			return 0
		}
	case scissors:
		if you == rock {
			return 6
		}
		if you == paper {
			return 0
		}
	}

	return 3
}

func main() {
	scan := bufio.NewScanner(strings.NewReader(input))

	total := 0
	for scan.Scan() {
		s := strings.Split(scan.Text(), " ")
		opp := opponent[s[0]]
		you := strategy2(opp, s[1])
		total += calcScore(opp, you)
	}

	println(total)

}

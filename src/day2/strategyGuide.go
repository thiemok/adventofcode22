package day2

import "strings"

var scores = map[string]int{
	"A X": 4,
	"A Y": 8,
	"A Z": 3,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 7,
	"C Y": 2,
	"C Z": 6,
}

var scoresRevised = map[string]int{
	"A X": 3,
	"A Y": 4,
	"A Z": 8,
	"B X": 1,
	"B Y": 5,
	"B Z": 9,
	"C X": 2,
	"C Y": 6,
	"C Z": 7,
}

func parseMoves(input string) []string {
	return strings.Split(input, "\n")
}

func CalculateGameScore(input string) int {
	moves := parseMoves(input)

	score := 0

	for _, v := range moves {
		score += scores[v]
	}

	return score
}

func CalculateRevisedGameScore(input string) int {
	moves := parseMoves(input)

	score := 0

	for _, v := range moves {
		score += scoresRevised[v]
	}

	return score
}

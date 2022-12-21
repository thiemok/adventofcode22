package day2

import (
	"fmt"
	"github.com/franela/goblin"
	"github.com/thiemok/adventofcode22/src/inputs"
	"testing"
)

func TestStrategyGuide(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("day2/strategyGuide", func() {
		input := `A Y
B X
C Z`

		g.It("correctly calculates the example score", func() {
			expected := 15

			g.Assert(CalculateGameScore(input)).Equal(expected)
		})

		g.It("correctly calculates the revised example score", func() {
			expected := 12

			g.Assert(CalculateRevisedGameScore(input)).Equal(expected)
		})
	})
}

func TestSolveDay2(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("day2/solution", func() {
		g.It("part one", func() {
			fmt.Println(CalculateGameScore(inputs.Day2))
		})
		g.It("part two", func() {
			fmt.Println(CalculateRevisedGameScore(inputs.Day2))
		})
	})
}

package day2

import (
	"fmt"
	"github.com/franela/goblin"
	"github.com/thiemok/adventofcode22/src/inputs"
	"testing"
)

func TestCalculateGameScore(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("day2/strategyGuid", func() {
		g.It("correctly calculates the example score", func() {
			input := `A Y
B X
C Z`
			expected := 15

			g.Assert(CalculateGameScore(input)).Equal(expected)
		})
	})
}

func TestSolveDay2(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("day2/solution", func() {
		g.It("part one", func() {
			fmt.Println(CalculateGameScore(inputs.Day2))
		})
	})
}

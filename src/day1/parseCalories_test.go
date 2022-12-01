package day1

import (
	"fmt"
	. "github.com/franela/goblin"
	"github.com/thiemok/adventofcode22/src/inputs"
	"sort"
	"testing"
)

func TestParseCalories(t *testing.T) {
	g := Goblin(t)
	g.Describe("day1/parseCalories", func() {
		input := `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`

		g.It("correctly parses total calories", func() {
			actual := ParseCalories(input)

			g.Assert(actual[0].Total).Equal(6000)
			g.Assert(actual[1].Total).Equal(4000)
			g.Assert(actual[2].Total).Equal(11000)
			g.Assert(actual[3].Total).Equal(24000)
			g.Assert(actual[4].Total).Equal(10000)
		})

		g.It("correctly parses individual food items", func() {
			actual := ParseCalories(input)

			g.Assert(actual[0].Items).Equal([]int{1000, 2000, 3000})
			g.Assert(actual[1].Items).Equal([]int{4000})
			g.Assert(actual[2].Items).Equal([]int{5000, 6000})
			g.Assert(actual[3].Items).Equal([]int{7000, 8000, 9000})
			g.Assert(actual[4].Items).Equal([]int{10000})
		})
	})
}

func TestSolveDay1(t *testing.T) {
	g := Goblin(t)
	g.Describe("day1/solution", func() {
		parsed := ParseCalories(inputs.Day1)
		sort.Slice(
			parsed,
			func(i, j int) bool {
				return parsed[i].Total > parsed[j].Total
		})

		g.It("part one", func() {
			fmt.Println(parsed[0].Total)
		})

		g.It("part two", func() {
			fmt.Println(parsed[0].Total + parsed[1].Total + parsed[2].Total)
		})
	})
}

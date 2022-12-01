package day1

import (
	"regexp"
	"strconv"
)

type CaloriesResult struct {
	Total int
	Items []int
	Id int
}

func ParseCalories(input string) []CaloriesResult {
	elves := splitElves(input)
	result := make([]CaloriesResult, len(elves))

	for i, v := range elves {
		items := parseItems(splitItems(v))
		total := calculateTotal(items)

		result[i] = CaloriesResult{Id: i, Items: items, Total: total}
	}

	return result
}

func splitElves(input string) []string {
	t := regexp.MustCompile(`\n\n`)
	return t.Split(input, -1)
}

func splitItems(input string) []string {
	t := regexp.MustCompile(`\n`)
	return t.Split(input, -1)
}

func parseItems(input []string) []int {
	items := make([]int, len(input))
	for i, v := range input {
		item, _ := strconv.Atoi(v)
		items[i] = item
	}

	return items
}

func calculateTotal(items []int) int {
	total := 0
	for _, v := range items {
		total += v
	}

	return total
}

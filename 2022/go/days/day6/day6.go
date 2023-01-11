package day6

import (
	"aoc/utils"
	"bufio"
	"fmt"
	"io"
)

func PartOne(s string) int {
	return firstSequenceOfDistinctChars(s, 4)
}

func PartTwo(s string) int {
	return firstSequenceOfDistinctChars(s, 14)
}

func firstSequenceOfDistinctChars(s string, distinctCount int) int {
	// create a new Counts struct
	counts := utils.CreateCounts()

	// iterate over the string, adding each rune to the counts
	for i, r := range s {
		counts.Add(r)
		if i >= distinctCount {
			counts.Remove(rune(s[i-distinctCount]))
		}

		// if the distinct count is 4, return the position of the last character
		if counts.Distinct() == distinctCount {
			return i + 1
		}
	}

	// if no sequence of 4 characters is found, return None
	return -1
}

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	data := scanner.Text()

	if output := PartOne(data); output != 1965 {
		panic(fmt.Errorf("PartOneDay6 failed -> %d", output))
	}

	if output := PartTwo(data); output != 2773 {
		panic(fmt.Errorf("PartTwoDay6 failed -> %d", output))
	}
}

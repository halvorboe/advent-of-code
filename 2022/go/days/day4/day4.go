package day4

import (
	"aoc/utils"
	"bufio"
	"io"
	"strconv"
	"strings"
)

func PartOne(data [][]utils.Range) int {
	count := 0
	for _, r := range data {
		firstRange, secondRange := r[0], r[1]
		if (firstRange.Start <= secondRange.Start && firstRange.End >= secondRange.End) || (secondRange.Start <= firstRange.Start && secondRange.End >= firstRange.End) {
			count++
		}
	}
	return count
}

func PartTwo(data [][]utils.Range) int {
	count := 0
	for _, r := range data {
		firstRange, secondRange := r[0], r[1]
		if (firstRange.Start <= secondRange.Start && firstRange.End >= secondRange.Start) || (secondRange.Start <= firstRange.Start && secondRange.End >= firstRange.Start) {
			count++
		}
	}
	return count
}

func parseRange(line string) []utils.Range {
	ranges := make([]utils.Range, 2)
	for i, r := range strings.Split(line, ",") {
		startEnd := strings.Split(r, "-")
		start, _ := strconv.Atoi(startEnd[0])
		end, _ := strconv.Atoi(startEnd[1])
		ranges[i] = utils.Range{Start: start, End: end}
	}
	return ranges
}

func Solve(reader io.Reader) {
	data := make([][]utils.Range, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data = append(data, parseRange(scanner.Text()))
	}
	if PartOne(data) != 582 {
		panic("PartOneDay4 failed")
	}
	if PartTwo(data) != 893 {
		panic("PartTwoDay4 failed")
	}
}

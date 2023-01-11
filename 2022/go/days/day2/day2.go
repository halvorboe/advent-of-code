package day2

import (
	"bufio"
	"io"
	"strings"
)

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, strings.TrimSpace(line))
	}
	if PartOne(lines) != 11873 {
		panic("part1 failed")
	}
	if PartTwo(lines) != 12014 {
		panic("part2 failed")
	}
}

func PartOne(data []string) int {
	total := 0
	for _, o := range data {
		switch o {
		case "A X":
			total += 1 + 3
		case "A Y":
			total += 2 + 6
		case "A Z":
			total += 3 + 0
		case "B X":
			total += 1 + 0
		case "B Y":
			total += 2 + 3
		case "B Z":
			total += 3 + 6
		case "C X":
			total += 1 + 6
		case "C Y":
			total += 2 + 0
		case "C Z":
			total += 3 + 3
		}
	}
	return total
}

func PartTwo(data []string) int {
	total := 0
	for _, o := range data {
		switch o {
		case "A X":
			total += 3 + 0
		case "A Y":
			total += 1 + 3
		case "A Z":
			total += 2 + 6
		case "B X":
			total += 1 + 0
		case "B Y":
			total += 2 + 3
		case "B Z":
			total += 3 + 6
		case "C X":
			total += 2 + 0
		case "C Y":
			total += 3 + 3
		case "C Z":
			total += 1 + 6
		}
	}
	return total
}

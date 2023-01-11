package day3

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

	if PartOne(lines) != 8349 {
		panic("part1 failed")
	}

	if PartTwo(lines) != 2681 {
		panic("part2 failed")
	}
}

func PartOne(data []string) int {
	r := make([]string, 0)
	for _, l := range data {
		half := len(l) / 2
		first := l[:half]
		second := l[half:]
		shared := intersect(first, second)
		r = append(r, shared...)
	}
	total := 0
	for _, c := range r {
		total += charToValue(c)
	}
	return total
}

func PartTwo(data []string) int {
	s := make([]string, 0)
	for i := 0; i < len(data); i += 3 {
		a := data[i]
		b := data[i+1]
		c := data[i+2]
		common := intersect(a, b, c)
		s = append(s, common...)
	}
	total := 0
	for _, c := range s {
		total += charToValue(c)
	}
	return total
}

func charToValue(c string) int {
	if c >= "a" && c <= "z" {
		return int(c[0]) - int('a') + 1
	} else if c >= "A" && c <= "Z" {
		return int(c[0]) - int('A') + 27
	}
	return 0
}

func intersect(strs ...string) []string {
	result := make([]string, 0)
	if len(strs) == 0 {
		return result
	}

	const charCountsSize = 128
	charCounts := make([]int, charCountsSize)

	for i := 0; i < len(strs); i++ {
		for _, c := range strs[i] {
			if charCounts[c] == i {
				charCounts[c]++
			}
		}
	}

	for i, count := range charCounts {
		if count == len(strs) {
			result = append(result, string(rune(i)))
		}
	}
	return result
}

package main

import (
	"aoc/days/day1"
	"aoc/days/day10"
	"aoc/days/day11"
	"aoc/days/day13"
	"aoc/days/day16"
	"aoc/days/day17"
	"aoc/days/day19"
	"aoc/days/day2"
	"aoc/days/day21"
	"aoc/days/day24"
	"aoc/days/day25"
	"aoc/days/day3"
	"aoc/days/day4"
	"aoc/days/day5"
	"aoc/days/day6"
	"aoc/days/day7"
	"aoc/days/day8"
	"aoc/days/day9"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

var INPUTS = []string{
	"inputs/day1_input.txt",
	"inputs/day2_input.txt",
	"inputs/day3_input.txt",
	"inputs/day4_input.txt",
	"inputs/day5_input.txt",
	"inputs/day6_input.txt",
	"inputs/day7_input.txt",
	"inputs/day8_input.txt",
	"inputs/day9_input.txt",
	"inputs/day10_input.txt",
	"inputs/day11_input.txt",
	"inputs/blank.txt",
	"inputs/blank.txt",
	"inputs/blank.txt",
	"inputs/blank.txt",
	"inputs/day16_input.txt",
	"inputs/day17_test.txt",
	"inputs/blank.txt",
	"inputs/day19_input.txt",
	"inputs/blank.txt",
	"inputs/day21_input.txt",
	"inputs/blank.txt",
	"inputs/blank.txt",
	"inputs/day24_input.txt",
	"inputs/day25_input.txt",
}
var FUNCTIONS = []func(io.Reader){
	day1.Solve,
	day2.Solve,
	day3.Solve,
	day4.Solve,
	day5.Solve,
	day6.Solve,
	day7.Solve,
	day8.Solve,
	day9.Solve,
	day10.Solve,
	day11.Solve,
	NotImplemented,
	day13.Solve,
	NotImplemented,
	NotImplemented,
	day16.Solve,
	day17.Solve,
	NotImplemented,
	day19.Solve,
	NotImplemented,
	day21.Solve,
	NotImplemented,
	NotImplemented,
	day24.Solve,
	day25.Solve,
}

func main() {

	inputStrings := make([]string, len(INPUTS))
	for i, input := range INPUTS {
		s, err := os.ReadFile(input)
		if err != nil {
			panic(err)
		}
		inputStrings[i] = string(s)
	}

	if len(os.Args) > 1 {
		day, _ := strconv.Atoi(os.Args[1])
		startTime := time.Now()
		FUNCTIONS[day-1](strings.NewReader(inputStrings[day-1]))
		endTime := time.Now()
		fmt.Printf("Day %d: %s\n", day, endTime.Sub(startTime))
		return
	}

	allStartTime := time.Now()

	for i, s := range inputStrings {
		startTime := time.Now()

		runDay(i, FUNCTIONS[i], s)
		endTime := time.Now()
		fmt.Printf("Day %d: %s\n", i+1, endTime.Sub(startTime))

	}
	allEndTime := time.Now()
	fmt.Printf("All: %s\n", time.Duration(allEndTime.Sub(allStartTime).Nanoseconds()))

}

func runDay(i int, f func(io.Reader), s string) {
	reader := strings.NewReader(s)
	f(reader)
}

func NotImplemented(io.Reader) {
	// fmt.Println("Not implemented")
}

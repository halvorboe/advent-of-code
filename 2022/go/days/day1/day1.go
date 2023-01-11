package day1

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func PartTwo(allSums []int) int {
	return allSums[len(allSums)-1] + allSums[len(allSums)-2] + allSums[len(allSums)-3]

}

func PartOne(allSums []int) int {
	return allSums[len(allSums)-1]
}

func Solve(reader io.Reader) {
	var allSums []int = make([]int, 0, 10000)
	var currentSum int = 0

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 1 {
			allSums = append(allSums, currentSum)
			currentSum = 0
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		currentSum += number

	}

	sort.Ints(allSums)

	if output := PartOne(allSums); output != 71506 { //  152 223971851179174 {
		panic(fmt.Errorf("PartOneDayTwentyOne failed -> %d", output))
	}

	if output := PartTwo(allSums); output != 209603 { // 301 3379022190351
		panic(fmt.Errorf("PartTwoDayTwentyOne failed -> %d", output))
	}

}

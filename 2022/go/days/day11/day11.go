package day11

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3
*/

const EMPTY = math.MinInt

type Operation int

const (
	ADD Operation = iota
	MUL
	ADD_SELF
	MUL_SELF
)

type ID uint8

type Monkey struct {
	id        ID
	operation Operation
	value     int
	divisor   int
	ifTrue    ID
	ifFalse   ID
}

const NONE = -1
const NUM_MONKEYS = 8

func PartOne(monkeys []Monkey, startingItems [][]int) int {
	return Simulate(monkeys, startingItems, 20, 3, true)
}

func PartTwo(monkeys []Monkey, startingItems [][]int) int {
	mod := 1
	for _, monkey := range monkeys {
		mod *= monkey.divisor
	}
	return Simulate(monkeys, startingItems, 10000, mod, false)
}

func Simulate(monkeys []Monkey, startingItems [][]int, rounds int, reducer int, useDivElseMod bool) int {
	var operationCount [NUM_MONKEYS]int

	for i, monkey := range monkeys {
		for _, item := range startingItems[i] {
			r := 0
			currentMonkey := &monkey
			currentItem := item
			for r < rounds {
				operationCount[currentMonkey.id]++
				// calculate new value
				switch currentMonkey.operation {
				case ADD:
					currentItem += currentMonkey.value
				case MUL:
					currentItem *= currentMonkey.value
				case ADD_SELF:
					currentItem += currentItem
				case MUL_SELF:
					currentItem *= currentItem
				}

				if useDivElseMod {
					currentItem /= reducer
				} else {
					currentItem %= reducer
				}
				// find next monkey
				if (currentItem % currentMonkey.divisor) == 0 {
					if currentMonkey.ifTrue < currentMonkey.id {
						r++
					}
					currentMonkey = &monkeys[currentMonkey.ifTrue]
				} else {
					if currentMonkey.ifFalse < currentMonkey.id {
						r++
					}
					currentMonkey = &monkeys[currentMonkey.ifFalse]
				}
			}
		}
	}
	sort.Ints(operationCount[:])
	return operationCount[len(operationCount)-1] * operationCount[len(operationCount)-2]
}

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	monkeys := make([]Monkey, NUM_MONKEYS)
	startingItems := make([][]int, NUM_MONKEYS)
	for i := range startingItems {
		startingItems[i] = make([]int, 0)
	}
	currentMonkey := ID(0)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			currentMonkey++
			continue
		}

		switch line[8] {
		case ':':
			monkeys[currentMonkey].id = currentMonkey
			continue
		case 'n':
			arr := strings.Split(line[18:], ", ")
			for _, item := range arr {
				i, _ := strconv.Atoi(item)
				startingItems[currentMonkey] = append(startingItems[currentMonkey], i)
			}
		case 'i':
			switch line[23] {
			case '+':
				monkeys[currentMonkey].operation = ADD
			case '*':
				monkeys[currentMonkey].operation = MUL
			}
			if line[25] == 'o' {
				monkeys[currentMonkey].value = EMPTY
				switch monkeys[currentMonkey].operation {
				case ADD:
					monkeys[currentMonkey].operation = ADD_SELF
				case MUL:
					monkeys[currentMonkey].operation = MUL_SELF
				}
				continue
			}
			value, err := strconv.Atoi(line[25:])
			if err != nil {
				panic(err)
			}
			monkeys[currentMonkey].value = value
		case 'd':
			divisor, err := strconv.Atoi(line[21:])
			if err != nil {
				panic(err)
			}
			monkeys[currentMonkey].divisor = divisor

		case 'r':
			target, _ := strconv.Atoi(line[29:])
			monkeys[currentMonkey].ifTrue = ID(target)
		case 'a':
			target, _ := strconv.Atoi(line[30:])
			monkeys[currentMonkey].ifFalse = ID(target)
		}
	}

	if output := PartOne(monkeys, startingItems); output != 76728 {
		panic(fmt.Errorf("PartOneDayEleven failed -> %d", output))
	}

	if output := PartTwo(monkeys, startingItems); output != 21553910156 {
		panic(fmt.Errorf("PartTwoDayEleven failed -> %d", output))
	}
}

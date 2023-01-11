package day5

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type StackMove struct {
	From   int
	To     int
	Number int
}

type Stack []rune

func (s *Stack) Push(element rune) {
	*s = append(*s, element)
}

func (s *Stack) Pop() rune {
	popped := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popped
}

func PartOne(stacks []Stack, moves []StackMove) string {
	for _, move := range moves {
		for i := 0; i < move.Number; i++ {
			letter := stacks[move.From].Pop()
			stacks[move.To].Push(letter)
		}
	}

	var result []rune
	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}

	return string(result)
}

func PartTwo(stacks []Stack, moves []StackMove) string {
	for _, move := range moves {
		// Create a new slice containing the top 'number' elements of stack 'source'
		toStackMove := make([]rune, 0)

		// Pop the top 'number' elements from stack 'source'
		for i := 0; i < move.Number; i++ {
			toStackMove = append(toStackMove, stacks[move.From].Pop())
		}

		// Push the elements in 'toStackMove' onto stack 'target'
		for i := range toStackMove {
			stacks[move.To].Push(toStackMove[len(toStackMove)-1-i])
		}
	}

	var result []rune
	for _, stack := range stacks {
		result = append(result, stack[len(stack)-1])
	}

	return string(result)

}

func parse(line string) StackMove {
	fields := strings.Fields(line)
	from, _ := strconv.Atoi(fields[3])
	to, _ := strconv.Atoi(fields[5])
	number, _ := strconv.Atoi(fields[1])
	return StackMove{from - 1, to - 1, number}
}

func createStacks() []Stack {
	stacks := make([]Stack, 0)

	initialStacks := [][]rune{
		{'N', 'B', 'D', 'T', 'V', 'G', 'Z', 'J'},
		{'S', 'R', 'M', 'D', 'W', 'P', 'F'},
		{'V', 'C', 'R', 'S', 'Z'},
		{'R', 'T', 'J', 'Z', 'P', 'H', 'G'},
		{'T', 'C', 'J', 'N', 'D', 'Z', 'Q', 'F'},
		{'N', 'V', 'P', 'W', 'G', 'S', 'F', 'M'},
		{'G', 'C', 'V', 'B', 'P', 'Q'},
		{'Z', 'B', 'P', 'N'},
		{'W', 'P', 'J'},
	}

	for _, stack := range initialStacks {
		stacks = append(stacks, Stack(stack))
	}

	return stacks
}

func Solve(reader io.Reader) {
	data := make([]StackMove, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		data = append(data, parse(scanner.Text()))
	}

	if output := PartOne(createStacks(), data); output != "GFTNRBZPF" {
		panic("PartOneDay5 failed -> " + output)
	}

	if output := PartTwo(createStacks(), data); output != "VRQWPDSGP" {
		panic("PartTwoDay5 failed -> " + output)
	}
}

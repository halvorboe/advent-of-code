package day10

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

const EMPTY = math.MinInt

type Operand int

const (
	NOOP Operand = iota
	ADDX
)

type Instruction struct {
	operand Operand
	arg1    int
}

type StateMachine struct {
	ip           int
	instructions []Instruction
	reg1         int
	reg2         int
}

func (sm *StateMachine) Tick() {
	if sm.reg2 != EMPTY {
		sm.reg1 = sm.reg2
		sm.reg2 = EMPTY
		sm.ip++
		return
	}
	instruction := sm.instructions[sm.ip]
	switch instruction.operand {
	case NOOP:
		sm.ip++
	case ADDX:
		sm.reg2 = sm.reg1 + instruction.arg1
	}
}

func createStateMachine(instructions []Instruction) StateMachine {
	return StateMachine{0, instructions, 1, EMPTY}
}

func PartOne(instructions []Instruction) int {
	sm := createStateMachine(instructions)
	t := 1
	s := 0
	for sm.ip < len(instructions) {
		if (t-20)%40 == 0 {
			s += t * sm.reg1
		}
		sm.Tick()
		t++
	}
	return s
}

func PartTwo(instructions []Instruction) int {
	sm := createStateMachine(instructions)
	crt := make([]bool, 40*6)
	t := 0
	for sm.ip < len(instructions) {
		if sm.reg1 >= (t%40)-1 && sm.reg1 <= (t%40)+1 {
			crt[t] = true
		}
		sm.Tick()
		t++
	}
	// hash the output
	count := 0
	for i := 0; i < len(crt); i++ {
		if crt[i] {
			count += i
		}
	}
	return count
}

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	instructions := make([]Instruction, 0)
	for scanner.Scan() {
		line := scanner.Text()
		switch line[0] {
		case 'a':
			value, err := strconv.Atoi(line[5:])
			if err != nil {
				panic(err)
			}
			instructions = append(instructions, Instruction{ADDX, value})
		case 'n':
			instructions = append(instructions, Instruction{NOOP, 0})
		}

	}
	if output := PartOne(instructions); output != 13740 {
		panic(fmt.Errorf("PartOneDayTen failed -> %d", output))
	}

	// using "hash" of output to verify
	if output := PartTwo(instructions); output != 10682 {
		panic(fmt.Errorf("PartTwoDayTen failed -> %d", output))
	}
}

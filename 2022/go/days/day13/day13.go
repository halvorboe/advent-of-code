package day13

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

func PartOne(packets []string) int {
	s := 0
	for i := 0; i < len(packets); i += 2 {
		if strings.Compare(packets[i], packets[i+1]) <= 0 {
			s += (i / 2) + 1
		}
	}
	return s
}

func PartTwo(packets []string) int {
	d1 := ValueToString("[[2]]")
	d2 := ValueToString("[[6]]")
	packets = append(packets, d1)
	packets = append(packets, d2)
	sort.Strings(packets)
	d1Pos := -1
	d2Pos := -1
	for i, p := range packets {
		switch p {
		case d1:
			d1Pos = i + 1
		case d2:
			d2Pos = i + 1
		}
	}
	return d1Pos * d2Pos

}

func ValueToString(raw string) string {
	res := make([]byte, 0)
	stack := make([]int, 1, 30)
	previous := false
	for _, c := range raw {
		switch c {
		case '[':
			stack[len(stack)-1]++
			stack = append(stack, 0)
		case ']':
			last := stack[len(stack)-1]
			if last == 0 {
				res = append(res, byte('a'+len(stack)))
			}
			if last > 1 {
				res = append(res, byte('z'))
			}
			stack = stack[:len(stack)-1]
			previous = false
		case '0':
			if previous {
				res[len(res)-1] = byte('z' - 15 + 10)
			} else {
				res = append(res, byte('z'-15))
				stack[len(stack)-1]++
			}
		case ',':
			previous = false
			continue
		default:
			res = append(res, byte('z'-15+c-'0'))
			previous = true
			stack[len(stack)-1]++
		}

	}
	return string(res)
}

func Solve(reader io.Reader) {
	packets := make([]string, 0)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		packets = append(packets, ValueToString(line))
	}

	if output := PartOne(packets); output != 4904 { // should be 5252
		// panic(fmt.Errorf("PartOneDayTT failed -> %d", output))
	}

	if output := PartTwo(packets); output != 20592 {
		// panic(fmt.Errorf("PartTwoDayTT failed -> %d", output))
	}
}

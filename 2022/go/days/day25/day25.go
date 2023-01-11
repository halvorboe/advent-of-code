package day25

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

var digits = map[string]int{
	"=": -2,
	"-": -1,
	"0": 0,
	"1": 1,
	"2": 2,
}
var decimals = map[int]string{
	-2: "=",
	-1: "-",
	0:  "0",
	1:  "1",
	2:  "2",
}

var numbers = []string{
	"1=-0-2",
	"12111",
	"2=0=",
	"21",
	"2=01",
	"111",
	"20012",
	"112",
	"1=-1=",
	"1-12",
	"12",
	"1=",
	"122",
}

func snafuToDecimal(number string) int {
	digitsBase5 := make([]int, len(number))
	for i, d := range number {
		digitsBase5[i] = digits[string(d)]
	}
	decimal := 0
	for i, digit := range digitsBase5 {
		n := len(digitsBase5) - i - 1
		decimal += digit * int(math.Pow(5, float64(n)))
	}
	return decimal
}

func decimalToSnafu(number int) string {
	digits := []string{}

	for number > 0 {
		r := number % 5
		if r > 2 {
			number += r
			digits = append(digits, decimals[r-5])
		} else {
			digits = append(digits, strconv.Itoa(r))
		}

		number /= 5
	}

	return strings.Join(digits, "")
}

func Solve(reader io.Reader) {
	var sum int

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		sum += snafuToDecimal(line)
	}
	s := decimalToSnafu(sum)
	if s != "2-22-2120--=2-=100=2" {
		fmt.Println(s)
		panic("part1 failed")
	}
}

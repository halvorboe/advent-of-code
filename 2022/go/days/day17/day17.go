package day17

import (
	"bufio"
	"fmt"
	"io"
)

const WIDTH = 7

type WindDirection rune

const (
	Left  WindDirection = '<'
	Right WindDirection = '>'
)

// if edge and piece then collide with edge
const EDGE uint = 0b100000001
const BLOCKED uint = 0b011111110 | EDGE

// pieces as bits with padding for edges
var PIECES = [][]uint{
	{
		0b000111100, // ####
	},
	{
		0b000010000, //  #
		0b000111000, // ###
		0b000010000, //  #
	},
	{
		0b000111000, // ###
		0b000001000, //   #
		0b000001000, //   #
	},
	{
		0b000100000, // #
		0b000100000, // #
		0b000100000, // #
		0b000100000, // #
	},
	{
		0b000110000, // ##
		0b000110000, // ##
	},
}

var COLUMNS = [WIDTH]uint{
	0b010000000,
	0b001000000,
	0b000100000,
	0b000010000,
	0b000001000,
	0b000000100,
	0b000000010,
}

func Fall(directions []WindDirection, n int) int {
	// fmt.Println(directions)
	t := 0
	d := 0
	rows := make([]uint, 0, 10000)
	rows = append(rows, BLOCKED)
	for t < n {
		piece := PIECES[t%len(PIECES)]

		row := len(rows) + 2 + len(piece)

		// fmt.Printf("t=%d, row=%d, piece=%v, len(rows)=%d\n", t, row, piece, len(rows))

		for {
			// shift piece left or right
			piece = shift(piece, directions[d])

			// collide piece
			if collide(rows, piece, row) {
				// fmt.Print("back -> ")
				// shift back
				shift(piece, opposite(directions[d]))
			}

			d = (d + 1) % len(directions)

			row--

			if collide(rows, piece, row) {
				// place piece
				// merge existing rows
				rows = merge(rows, piece, row+3)
				break
			}
			// fmt.Println("down")

		}
		t++
		// printRows(rows)
	}
	return len(rows)

}

func opposite(direction WindDirection) WindDirection {
	if direction == Left {
		return Right
	}
	return Left
}

func shift(piece []uint, direction WindDirection) []uint {
	if direction == Left {
		// fmt.Println("left")
		for i := range piece {
			piece[i] = piece[i] << 1
		}
	} else {
		// fmt.Println("right")
		for i := range piece {
			piece[i] = piece[i] >> 1
		}
	}
	return piece
}

func merge(rows []uint, piece []uint, offset int) []uint {
	// fmt.Println(rows, piece, offset)
	for i := 0; i < len(piece); i++ {
		// row
		// 5 - (3 - 0) = 2
		// 5 - (3 - 1) = 3
		// 5 - (3 - 2) = 4
		j := offset - i
		if j >= len(rows) {
			// fmt.Printf("append %09b\n", piece[i]|EDGE)
			rows = append(rows, piece[i]|EDGE)
		} else {
			// fmt.Printf("merge %09b %09b\n", rows[j], piece[i])
			rows[j] |= piece[i]
		}
	}
	return rows
}

// could be optimized based on direction and piece
func collide(rows []uint, piece []uint, offset int) bool {

	for i := len(piece) - 1; i >= 0; i-- {
		// fmt.Printf("i=%d, offset=%d, len(rows)=%d\n", i, offset, len(rows))
		if len(rows) > offset-i {
			// fmt.Println(rows, offset, i)
			// fmt.Printf("%09b\n", rows[offset-i])
			// fmt.Printf("%09b\n", piece[i])
			row := rows[offset-i]
			if row&piece[i] != 0 {
				return true
			}
		} else {
			if piece[i]&EDGE != 0 {
				return true
			}
		}
	}
	return false
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func PartOne(directions []WindDirection) int {
	return Fall(directions, 2022)
}

func PartTwo(directions []WindDirection) int {
	return 0
}

func Solve(reader io.Reader) {
	packets := make([]WindDirection, 0, 2000)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		for _, c := range line {
			packets = append(packets, WindDirection(c))
		}
	}

	if output := PartOne(packets); output != 3361 { // should be 5252
		// panic(fmt.Errorf("PartOneDay17 failed -> %d", output))
	}

	if output := PartTwo(packets); output != 0 {
		// panic(fmt.Errorf("PartTwoDay17 failed -> %d", output))
	}
}

func printRows(rows []uint) {
	fmt.Println("--rows--")
	for i := len(rows) - 1; i >= 0; i-- {
		fmt.Printf("%09b\n", rows[i])
	}
}

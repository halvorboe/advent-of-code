package day9

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

const SET_SIZE = 600
const OFFSET = SET_SIZE / 2

type Move struct {
	dx    int
	dy    int
	steps int
}

type Point struct {
	x int
	y int
}

func PartOne(moves []Move) int {
	return SimulateRopeOfLength(moves, 2)

}

func PartTwo(moves []Move) int {
	return SimulateRopeOfLength(moves, 10)
}

func SimulateRopeOfLength(moves []Move, length int) int {
	rope := make([]Point, length)
	visible := make([][SET_SIZE]bool, SET_SIZE)
	count := 0

	for _, move := range moves {
		for s := int(0); s < move.steps; s++ {
			rope[0].x += move.dx
			rope[0].y += move.dy

			for i := 1; i < len(rope); i++ {

				if DistanceLessThanTwo(rope[i], rope[i-1]) {
					break
				}

				trailingPoint := &rope[i]
				leadingPoint := &rope[i-1]

				if trailingPoint.x < leadingPoint.x {
					trailingPoint.x++
				} else if trailingPoint.x > leadingPoint.x {
					trailingPoint.x--
				}
				if trailingPoint.y < leadingPoint.y {
					trailingPoint.y++
				} else if trailingPoint.y > leadingPoint.y {
					trailingPoint.y--
				}
			}
			x, y := rope[len(rope)-1].x+OFFSET, rope[len(rope)-1].y+OFFSET
			if !visible[x][y] {
				visible[x][y] = true
				count += 1
			}
		}
	}

	return count
}

func DistanceLessThanTwo(trailingPoint, leadingPoint Point) bool {
	var x, y int
	if trailingPoint.x == leadingPoint.x {
		x = 0
	} else if trailingPoint.x > leadingPoint.x {
		x = trailingPoint.x - leadingPoint.x
	} else {
		x = leadingPoint.x - trailingPoint.x
	}

	if trailingPoint.y == leadingPoint.y {
		y = 0
	} else if trailingPoint.y > leadingPoint.y {
		y = trailingPoint.y - leadingPoint.y
	} else {
		y = leadingPoint.y - trailingPoint.y
	}
	return x*x+y*y <= 2
}

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	moves := make([]Move, 0)
	for scanner.Scan() {
		line := scanner.Text()
		move := Move{}
		switch line[0] {
		case 'U':
			move.dx = 0
			move.dy = -1
		case 'D':
			move.dx = 0
			move.dy = 1
		case 'L':
			move.dx = -1
			move.dy = 0
		case 'R':
			move.dx = 1
			move.dy = 0
		}
		steps, _ := strconv.Atoi(line[2:])
		move.steps = steps
		moves = append(moves, move)
	}
	if output := PartOne(moves); output != 6384 {
		panic(fmt.Errorf("PartOneDayNine failed -> %d", output))
	}

	if output := PartTwo(moves); output != 2734 {
		panic(fmt.Errorf("PartTwoDayNine failed -> %d", output))
	}
}

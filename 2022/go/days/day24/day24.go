package day24

import (
	"bufio"
	"fmt"
	"io"
	"math"
)

var POSITIONS = []position{{-1, 0}, {0, -1}, {0, 0}, {0, 1}, {1, 0}}

type position struct {
	x int
	y int
}

type positionTime struct {
	position   position
	time       int
	estimation int
}

func nMod(d, m int) int {
	res := d % m
	if res < 0 {
		return res + m
	}
	return res
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int) int {
	return a * b / GCD(a, b)
}

func search(storms *BoolSet, start, end position, time int) int {
	endX, endY := end.x, end.y
	lenX, lenY := storms.x-2, storms.y-2
	epoch := LCM(lenX, lenY)

	pruned := CreateBoolSet(lenX, lenY, epoch)
	// hasStorm := create(lenX, lenY, epoch)

	pruned.Set(start.x, start.y, time, true)
	queue := make(priorityQueue, 0, 2000)
	queue = append(queue, positionTime{start, time, time + abs(start.x-endX) + abs(start.y-endY)})

	for queue.Len() > 0 {
		item := HeapPop(&queue)

		x, y := item.position.x, item.position.y
		t := item.time + 1
		for _, move := range POSITIONS {
			x, y := x+move.x, y+move.y
			// out of bounds
			// start
			if x == start.x && y == start.y {
				// do nothing
			} else if x == end.x && y == end.y {
				return t
			} else if y < 1 || y >= storms.y-1 || x < 1 || x >= storms.x-1 {
				continue
			}

			// already pruned
			if pruned.Get(x, y, t%epoch) {
				continue
			}

			// if we either have a storm or have been in the position before at the same timing we skip.
			// it's safe to skip because if we're in the same spot with the same timing we can't get to the end faster.
			if storms.Get(nMod(x-1+t, lenX)+1, y, 0) {
				pruned.Set(x, y, t%epoch, true)
				continue
			} else if storms.Get(nMod(x-1-t, lenX)+1, y, 1) {
				pruned.Set(x, y, t%epoch, true)
				continue
			} else if storms.Get(x, nMod(y-1+t, lenY)+1, 2) {
				pruned.Set(x, y, t%epoch, true)
				continue
			} else if storms.Get(x, nMod(y-1-t, lenY)+1, 3) {
				pruned.Set(x, y, t%epoch, true)
				continue
			} else {
				pruned.Set(x, y, t%epoch, true)
				HeapPush(&queue, positionTime{position{x, y}, t, t + abs(x-endX) + abs(y-endY)})
			}
		}
	}

	return math.MaxInt32
}

func Parts(lines *BoolSet) (int, int) {
	start := position{1, 0}
	end := position{lines.x - 2, lines.y - 1}
	p1 := search(lines, start, end, 0)
	p2 := search(
		lines, start, end,
		search(lines, end, start, p1),
	)
	return p1, p2
}

func Solve(reader io.Reader) {
	lines := make([][]rune, 0)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		runes := make([]rune, 0)
		for _, c := range line {
			runes = append(runes, c)
		}
		lines = append(lines, runes)
	}

	storms := CreateBoolSet(len(lines[0]), len(lines), 4)

	for y, line := range lines {
		for x, c := range line {
			switch c {
			case '<':
				storms.Set(x, y, 0, true)
			case '>':
				storms.Set(x, y, 1, true)
			case '^':
				storms.Set(x, y, 2, true)
			case 'v':
				storms.Set(x, y, 3, true)
			}
		}
	}

	p1, p2 := Parts(&storms)

	if p1 != 274 { //  152 223971851179174 {
		panic(fmt.Errorf("PartOneDayTwentyOne failed -> %d", p1))
	}

	if p2 != 839 { // 301 3379022190351
		panic(fmt.Errorf("PartTwoDayTwentyOne failed -> %d", p2))
	}

}

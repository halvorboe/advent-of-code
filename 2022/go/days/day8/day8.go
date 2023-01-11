package day8

import (
	"bufio"
	"fmt"
	"io"
)

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

func PartOne(trees [][]int) int {
	visible := make([][]bool, len(trees))
	for i := range visible {
		visible[i] = make([]bool, len(trees[0]))
	}

	for y := 0; y < len(trees[0]); y++ {
		// down
		current := -1
		for i := 0; i < len(trees); i++ {
			tree := trees[i][y]
			if tree > current {
				current = tree
				visible[i][y] = true
			}
		}

		// up
		current = -1
		for i := len(trees) - 1; i >= 0; i-- {
			tree := trees[i][y]
			if tree > current {
				current = tree
				visible[i][y] = true
			}
		}
	}

	for x := 0; x < len(trees); x++ {
		// right
		current := -1
		for i := 0; i < len(trees[0]); i++ {
			tree := trees[x][i]
			if tree > current {
				current = tree
				visible[x][i] = true
			}
		}

		// left
		current = -1
		for i := len(trees[0]) - 1; i >= 0; i-- {
			tree := trees[x][i]
			if tree > current {
				current = tree
				visible[x][i] = true
			}
		}
	}

	count := 0
	for _, row := range visible {
		for _, v := range row {
			if v {
				count++
			}
		}
	}
	return count

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func PartTwo(trees [][]int) int {
	maxViewingDistance := 0
	for x := 1; x < len(trees)-1; x++ {
		for y := 1; y < len(trees[0])-1; y++ {
			height := trees[x][y]

			// don't consider trees that are too short
			if height <= 5 {
				continue
			}

			up := 0
			for i := x - 1; i >= 0; i-- {
				tree := trees[i][y]
				up++
				if tree >= height {
					break
				}
			}

			down := 0
			for i := x + 1; i < len(trees); i++ {
				tree := trees[i][y]
				down++
				if tree >= height {
					break
				}
			}

			left := 0
			for i := y - 1; i >= 0; i-- {
				tree := trees[x][i]
				left++
				if tree >= height {
					break
				}
			}

			right := 0
			for i := y + 1; i < len(trees[0]); i++ {
				tree := trees[x][i]
				right++
				if tree >= height {
					break
				}
			}

			maxViewingDistance = max(maxViewingDistance, up*down*left*right)
		}
	}

	return maxViewingDistance
}

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	trees := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		treeLine := make([]int, 0)
		for _, c := range line {
			treeLine = append(treeLine, int(c)-48)
		}
		trees = append(trees, treeLine)
	}

	if output := PartOne(trees); output != 1690 {
		panic(fmt.Errorf("PartOneDayEight failed -> %d", output))
	}

	if output := PartTwo(trees); output != 535680 {
		panic(fmt.Errorf("PartTwoDayEight failed -> %d", output))
	}
}

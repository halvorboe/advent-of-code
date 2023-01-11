package day16

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// https://github.com/Crazytieguy/advent-of-code/blob/master/2022/src/bin/day16/main.rs

const INF = 1000000000

type Valve struct {
	From string
	To   []string
	Flow int
}

type Score struct {
	visited uint16
	score   int
}

type Node struct {
	id        string
	distances []int
	flow      int
}

type State struct {
	visited        uint16
	totalFlow      int
	remainingMoves int
	position       int
	cachedBound    int
}

func (state *State) String() string {
	return fmt.Sprintf("State{visited: %016b, totalFlow: %d, remainingMoves: %d, position: %d}", state.visited, state.totalFlow, state.remainingMoves, state.position)
}

// Get child nodes
func (state *State) Branch(states []State, nodes []Node) []State {
	states = states[:0]
	for i := 1; i < len(nodes); i++ {
		// starting node or already visited
		if i == state.position {
			continue
		}
		if state.visited&(1<<i) != 0 {
			continue
		}
		remainingMoves := state.remainingMoves - nodes[state.position].distances[i]
		if remainingMoves <= 0 {
			continue
		}
		branchState := State{
			visited:        state.visited | (1 << i),
			totalFlow:      state.totalFlow + (nodes[i].flow * remainingMoves),
			remainingMoves: remainingMoves,
			position:       i,
		}
		branchState.cacheBound(nodes)
		states = append(states, branchState)
	}
	return states
}

// Calculate final value based on visiting in decending order
func (state *State) cacheBound(graph []Node) {
	totalFlow := state.totalFlow
	remainingMoves := state.remainingMoves
	currentNode := state.position
	for nextNode := 1; nextNode < len(graph); nextNode++ {
		if state.visited&(1<<nextNode) != 0 {
			continue
		}
		remainingMoves -= graph[currentNode].distances[nextNode]
		if remainingMoves <= 0 {
			continue
		}
		currentNode = nextNode
		totalFlow += graph[currentNode].flow * remainingMoves
	}
	state.cachedBound = totalFlow
}

func (state *State) Bound() int {
	if state.cachedBound == 0 {
		panic("cachedBound not set")
	}
	return state.cachedBound
}

func BranchAndBound(nodes []Node, remainingMoves int) Uint16Map {
	// check if current is better than best
	queue := NewPartitionedQueue(remainingMoves)
	startState := State{
		visited:        0b1,
		totalFlow:      0,
		remainingMoves: remainingMoves,
		position:       0,
	}
	startState.cacheBound(nodes)
	queue.Push(startState)
	bestPerVisited := NewUint16Map()
	best := 0
	states := make([]State, 0, len(nodes))
	for queue.Len() > 0 {
		state := queue.Pop()
		if state.Bound() > best {
			best = state.Bound()
		}
		if state.Bound() > bestPerVisited[state.visited] {
			bestPerVisited[state.visited] = state.Bound()
		} else {
			continue
		}
		states = state.Branch(states, nodes)
		for _, branch := range states {
			if branch.Bound() > best/2 && branch.Bound() > bestPerVisited[branch.visited] {
				queue.Push(branch)
			}
		}
	}

	return bestPerVisited
}

func PartOne(nodes []Node) int {
	bestPerVisited := BranchAndBound(nodes, 30)

	best := 0
	for _, v := range bestPerVisited {
		if v > best {
			best = v
		}
	}
	return best
}

func PartTwo(nodes []Node) int {
	bestPerVisited := BranchAndBound(nodes, 26)

	scores := make([]Score, 0, len(bestPerVisited))
	for visited, score := range bestPerVisited {
		if score == 0 {
			continue
		}
		scores = append(scores, Score{
			visited: uint16(visited),
			score:   score,
		})
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i].score > scores[j].score
	})

	best := 0
	for i := 0; i < len(scores); i++ {
		for j := i + 1; j < len(scores); j++ {
			myScore := scores[i].score
			myVisited := scores[i].visited
			otherScore := scores[j].score
			otherVisited := scores[j].visited
			score := myScore + otherScore
			if score <= best {
				break
			}
			if myVisited&otherVisited == 1 {
				best = score
				break
			}

		}
	}

	return best
}

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	valves := make([]Valve, 0, 100)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		to := make([]string, 0)
		for _, t := range parts[9:] {
			to = append(to, strings.TrimSuffix(t, ","))
		}
		flow, err := strconv.Atoi(strings.TrimSuffix(parts[4][5:], ";"))
		if err != nil {
			panic(err)
		}
		valve := Valve{
			From: parts[1],
			To:   to,
			Flow: flow,
		}
		i++
		valves = append(valves, valve)
	}

	graph := buildGraph(valves)

	if output := PartOne(graph); output != 1737 {
		panic(fmt.Errorf("PartOneDay16 failed -> %d", output))
	}

	if output := PartTwo(graph); output != 2216 {
		panic(fmt.Errorf("PartTwoDay16 failed -> %d", output))
	}
}

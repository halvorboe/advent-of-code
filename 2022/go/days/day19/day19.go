package day19

import (
	"bufio"
	"fmt"
	"hash/maphash"
	"io"
)

const (
	ORE      = 0
	CLAY     = 1
	OBSIDIAN = 2
	GEODE    = 3
)

type Execution struct {
	robots     Amounts
	resources  Amounts
	minutes    int
	upperCache int
}

func (ex *Execution) Hash() uint64 {
	var h maphash.Hash
	h.WriteByte(byte(ex.robots[0]))
	h.WriteByte(byte(ex.robots[1]))
	h.WriteByte(byte(ex.robots[2]))
	h.WriteByte(byte(ex.robots[3]))
	h.WriteByte(byte(ex.resources[0]))
	h.WriteByte(byte(ex.resources[1]))
	h.WriteByte(byte(ex.resources[2]))
	h.WriteByte(byte(ex.resources[3]))
	h.WriteByte(byte(ex.minutes))
	return h.Sum64()
}

const EMPTY = -1

func (ex *Execution) Clone() Execution {
	return Execution{
		robots:     ex.robots.Clone(),
		resources:  ex.resources.Clone(),
		minutes:    ex.minutes,
		upperCache: EMPTY,
	}
}

func (ex *Execution) String() string {
	return fmt.Sprintf("Execution{robots: %v, resources: %v, minutes: %d}", ex.robots, ex.resources, ex.minutes)
}

// https://github.com/orlp/aoc2022/blob/master/src/bin/day19.rs

func BestNumGeodes(bp *Blueprint, minutes int) int {
	best := 0
	executions := createPriorityQueue(bp, minutes)
	executions.Push(&Execution{
		robots:     Amounts{1, 0, 0, 0},
		resources:  Amounts{0, 0, 0, 0},
		minutes:    0,
		upperCache: EMPTY,
	})
	seen := make(map[uint64]bool, 5000)
	for executions.Len() > 0 {
		ex := HeapPop(executions)
		if ex.upperGeodes(bp, minutes) < best {
			break
		}

		for _, resource := range RESOURCES {
			if next := ex.buildRobots(bp, resource, minutes); next != nil {
				upper := next.upperGeodes(bp, minutes)
				if upper > best {
					lower := next.lowerGeodes(bp, minutes)
					if lower > best {
						best = lower
					}
					hash := next.Hash()
					if next.minutes < minutes && !seen[hash] {
						seen[hash] = true
						HeapPush(executions, next)
					}
				}
			}
		}

	}

	return best
}

func (ex *Execution) buildRobots(bp *Blueprint, resource Resource, maxMins int) *Execution {
	var haveEnoughAlready bool
	var costs [3]int

	switch resource {
	case ORE:
		haveEnoughAlready = ex.robots[ORE] >= bp.MaxCost()
		costs = [3]int{bp.oreCosts[ORE], 0, 0}
	case CLAY:
		haveEnoughAlready = ex.robots[CLAY] >= bp.obsidianClayCost
		costs = [3]int{bp.oreCosts[CLAY], 0, 0}
	case OBSIDIAN:
		haveEnoughAlready = ex.robots[OBSIDIAN] >= bp.geodeObsidianCost
		costs = [3]int{bp.oreCosts[OBSIDIAN], bp.obsidianClayCost, 0}
	case GEODE:
		haveEnoughAlready = ex.robots[GEODE] >= bp.geodeObsidianCost
		costs = [3]int{bp.oreCosts[GEODE], 0, bp.geodeObsidianCost}
	default:
		panic("invalid resource")
	}

	delay := 0
	for _, r := range [3]int{ORE, CLAY, OBSIDIAN} {
		if costs[r] <= ex.resources[r] {
			continue
		}
		if ex.robots[r] == 0 {
			return nil
		}
		d := 1 + (costs[r]-ex.resources[r]+ex.robots[r]-1)/ex.robots[r]
		if d > delay {
			delay = d
		}
	}
	if delay == 0 {
		delay = 1
	}
	minutes := ex.minutes + delay
	if haveEnoughAlready || minutes > maxMins {
		return nil
	}

	ret := ex.Clone()
	for r := 0; r < 3; r++ {
		ret.resources[r] += delay*ret.robots[r] - costs[r]
	}
	ret.resources[GEODE] += ret.robots[GEODE] * delay

	ret.minutes += delay
	ret.robots[resource]++
	return &ret
}

func (ex *Execution) lowerGeodes(bp *Blueprint, maxMins int) int {
	// We only construct geode bots.
	robots := ex.robots.Clone()
	resources := ex.resources.Clone()
	for m := ex.minutes; m < maxMins; m++ {
		newBot := resources[ORE] >= bp.oreCosts[GEODE] && resources[OBSIDIAN] >= bp.geodeObsidianCost
		for r := 0; r < 4; r++ {
			resources[r] += robots[r]
		}
		if newBot {
			resources[ORE] -= bp.oreCosts[GEODE]
			resources[OBSIDIAN] -= bp.geodeObsidianCost
			robots[GEODE]++

		}
	}
	return resources[GEODE]
}

func (ex *Execution) upperGeodes(bp *Blueprint, maxMins int) int {
	// We greedily build robots, but the costs for one type of robot are
	// not subtracted from the pool of resources of the other robots, and we
	// can build multiple robot types at once.
	if ex.upperCache != EMPTY {
		return ex.upperCache
	}

	robots := ex.robots
	oresFor := Amounts{ex.resources[ORE]}
	clay, obs, geodes := ex.resources[CLAY], ex.resources[OBSIDIAN], ex.resources[GEODE]
	for m := ex.minutes; m <= maxMins; m++ {
		newBot := [4]bool{
			oresFor[ORE] >= bp.oreCosts[ORE],
			oresFor[CLAY] >= bp.oreCosts[CLAY],
			oresFor[OBSIDIAN] >= bp.oreCosts[OBSIDIAN] && clay >= bp.obsidianClayCost,
			oresFor[GEODE] >= bp.oreCosts[GEODE] && obs >= bp.geodeObsidianCost,
		}

		for r := 0; r < 4; r++ {
			oresFor[r] += robots[ORE]
			if newBot[r] {
				oresFor[r] -= bp.oreCosts[r]
			}
		}
		clay += robots[CLAY]
		if newBot[OBSIDIAN] {
			clay -= bp.obsidianClayCost
		}
		obs += robots[OBSIDIAN]
		if newBot[GEODE] {
			obs -= bp.geodeObsidianCost
		}
		geodes += robots[GEODE]

		for r := 0; r < 4; r++ {
			if newBot[r] {
				robots[r]++
			}
		}
	}

	ex.upperCache = geodes

	return ex.upperCache
}

func PartOne(blueprints []Blueprint) int {
	sum := 0
	for i, bp := range blueprints {
		best := BestNumGeodes(&bp, 24)
		sum += best * (i + 1)
	}
	return sum
}

func PartTwo(blueprints []Blueprint) int {
	return BestNumGeodes(&blueprints[0], 32) * BestNumGeodes(&blueprints[1], 32) * BestNumGeodes(&blueprints[2], 32)
}

func Solve(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	blueprints := make([]Blueprint, 0)
	for scanner.Scan() {
		line := scanner.Text()
		bp := ParseBlueprint(line)
		blueprints = append(blueprints, bp)
	}
	if output := PartOne(blueprints); output != 2301 {
		panic(fmt.Errorf("part 1: %d", output))
	}
	if output := PartTwo(blueprints); output != 10336 {
		panic(fmt.Errorf("part 2: %d", output))
	}

}

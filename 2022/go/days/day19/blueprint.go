package day19

import (
	"fmt"
	"strconv"
	"strings"
)

type Blueprint struct {
	oreCosts          Amounts
	obsidianClayCost  int
	geodeObsidianCost int
}

func ParseBlueprint(s string) Blueprint {
	// Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.
	fields := strings.Fields(s)
	if len(fields) < 10 {
		panic(fmt.Errorf("invalid input: %s", s))
	}

	bp := Blueprint{}

	ore, err := strconv.Atoi(fields[6])
	if err != nil {
		panic(err)
	}
	bp.oreCosts[Ore] = ore

	clay, err := strconv.Atoi(fields[12])
	if err != nil {
		panic(err)
	}
	bp.oreCosts[Clay] = clay

	obsOre, err := strconv.Atoi(fields[18])
	if err != nil {
		panic(err)
	}
	bp.oreCosts[Obsidian] = obsOre

	obs, err := strconv.Atoi(fields[21])
	if err != nil {
		panic(err)
	}
	bp.obsidianClayCost = obs

	geoOre, err := strconv.Atoi(fields[27])
	if err != nil {
		panic(err)
	}
	bp.oreCosts[Geode] = geoOre

	geo, err := strconv.Atoi(fields[30])
	if err != nil {
		panic(err)
	}
	bp.geodeObsidianCost = geo

	return bp
}

func (bp *Blueprint) MaxCost() int {
	return max(bp.oreCosts[:]...)
}

func max(nums ...int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

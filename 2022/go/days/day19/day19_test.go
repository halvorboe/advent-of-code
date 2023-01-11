package day19

import (
	"fmt"
	"testing"
)

var BP1 = ParseBlueprint("Blueprint 1: Each ore robot costs 4 ore. Each clay robot costs 2 ore. Each obsidian robot costs 3 ore and 14 clay. Each geode robot costs 2 ore and 7 obsidian.")
var BP2 = ParseBlueprint("Blueprint 2: Each ore robot costs 2 ore. Each clay robot costs 3 ore. Each obsidian robot costs 3 ore and 8 clay. Each geode robot costs 3 ore and 12 obsidian.")

var EX1 = Execution{
	robots:    Amounts{1, 0, 0, 0},
	resources: Amounts{0, 0, 0, 0},
	minutes:   0,
}
var EX2 = Execution{
	robots:    Amounts{0, 0, 0, 1},
	resources: Amounts{0, 0, 0, 0},
	minutes:   0,
}
var EX3 = Execution{
	robots:    Amounts{1, 1, 1, 0},
	resources: Amounts{0, 0, 0, 0},
	minutes:   0,
}
var EX4 = Execution{
	robots:    Amounts{3, 13, 6, 0},
	resources: Amounts{4, 1, 1, 0},
	minutes:   0,
}

func TestBest(t *testing.T) {
	if output := BestNumGeodes(&BP1, 1); output != 0 {
		t.Error("expected at least 1 geode", output)
	}
	if output := BestNumGeodes(&BP1, 10); output != 0 {
		t.Error("expected at least 1 geode", output)
	}
	if output := BestNumGeodes(&BP1, 24); output != 9 {
		t.Error("expected at least 1 geode", output)
	}
	fmt.Println("--BP2--")
	if output := BestNumGeodes(&BP2, 1); output != 0 {
		t.Error("expected at least 1 geode", output)
	}
	if output := BestNumGeodes(&BP2, 10); output != 0 {
		t.Error("expected at least 1 geode", output)
	}
	if output := BestNumGeodes(&BP2, 24); output != 12 {
		t.Error("expected at least 1 geode", output)
	}
	// longer
	if output := BestNumGeodes(&BP2, 32); output != 62 {
		t.Error("expected at least 1 geode", output)
	}
}

func TestLower(t *testing.T) {
	if output := EX1.lowerGeodes(&BP1, 24); output != 0 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX2.lowerGeodes(&BP1, 24); output != 24 {
		t.Error("expected at least 1 geode", output)
	}
}

func TestUpper(t *testing.T) {
	if output := EX1.upperGeodes(&BP1, 24); output != 22 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX2.lowerGeodes(&BP1, 24); output != 24 {
		t.Error("expected at least 1 geode", output)
	}
}

func TestBuild(t *testing.T) {
	// ore
	if output := EX1.buildRobots(&BP1, Ore, 24); output.minutes != 5 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX1.buildRobots(&BP2, Ore, 24); output.minutes != 3 {
		t.Error("expected at least 1 geode", output)
	}
	// clay
	if output := EX1.buildRobots(&BP1, Clay, 24); output.minutes != 3 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX1.buildRobots(&BP2, Clay, 24); output.minutes != 4 {
		t.Error("expected at least 1 geode", output)
	}
	// obsidian
	if output := EX1.buildRobots(&BP1, Obsidian, 24); output != nil {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX1.buildRobots(&BP2, Obsidian, 24); output != nil {
		t.Error("expected at least 1 geode", output)
	}

	// all starting with 1 of each
	if output := EX3.buildRobots(&BP1, Obsidian, 24); output.minutes != 15 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX3.buildRobots(&BP2, Obsidian, 24); output.minutes != 9 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX3.buildRobots(&BP1, Geode, 24); output.minutes != 8 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX3.buildRobots(&BP2, Geode, 24); output.minutes != 13 {
		t.Error("expected at least 1 geode", output)
	}

	// more complicated
	if output := EX4.buildRobots(&BP1, Ore, 24); output.minutes != 1 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX4.buildRobots(&BP2, Ore, 24); output != nil {
		t.Error("expected at least 1 geode", output)
	}

	if output := EX4.buildRobots(&BP1, Clay, 24); output.minutes != 1 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX4.buildRobots(&BP2, Clay, 24); output != nil {
		t.Error("expected at least 1 geode", output)
	}

	if output := EX4.buildRobots(&BP1, Obsidian, 24); output.minutes != 2 {
		t.Error("expected at least 1 geode", output)
	}
	if output := EX4.buildRobots(&BP2, Obsidian, 24); output.minutes != 2 {
		t.Error("expected at least 1 geode", output)
	}

}

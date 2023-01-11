package day19

type Resource int

const (
	Ore Resource = iota
	Clay
	Obsidian
	Geode
)

var RESOURCES = [4]Resource{Ore, Clay, Obsidian, Geode}

type Amounts [4]int

func (a *Amounts) Clone() Amounts {
	return Amounts{a[0], a[1], a[2], a[3]}
}

func (a *Amounts) Hash() uint64 {
	return uint64(a[0])<<48 | uint64(a[1])<<32 | uint64(a[2])<<16 | uint64(a[3])
}

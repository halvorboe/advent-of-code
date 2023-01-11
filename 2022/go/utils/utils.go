package utils

type Range struct {
	Start int
	End   int
}

type Counts struct {
	counts          []int
	distinctCounter int
}

func (c *Counts) Add(r rune) {
	if c.counts[r] == 0 {
		c.distinctCounter++
	}
	// increase the count of the rune
	c.counts[r]++
}

func (c *Counts) Remove(r rune) {
	if c.counts[r] == 1 {
		c.distinctCounter--
	}
	// decrement the count of the rune
	c.counts[r]--
}

func (c *Counts) Distinct() int {
	return c.distinctCounter
}

func CreateCounts() Counts {
	const charCountsSize = 256
	return Counts{make([]int, charCountsSize), 0}
}

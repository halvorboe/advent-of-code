package day24

type BoolSet struct {
	raw []bool
	x   int
	y   int
	z   int
}

func (bs *BoolSet) Set(x, y, z int, v bool) {
	bs.raw[bs.computeIndex(x, y, z)] = v
}

func (bs *BoolSet) Get(x, y, z int) bool {
	return bs.raw[bs.computeIndex(x, y, z)]
}

func (bs *BoolSet) computeIndex(x, y, z int) int {
	return x + y*bs.x + z*bs.x*bs.y
}

func CreateBoolSet(x, y, z int) BoolSet {
	return BoolSet{make([]bool, x*y*z), x, y, z}
}

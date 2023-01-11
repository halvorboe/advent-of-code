package day19

type priorityQueue struct {
	bp      *Blueprint
	pq      []*Execution
	minutes int
}

func createPriorityQueue(bp *Blueprint, minutes int) *priorityQueue {
	return &priorityQueue{bp, make([]*Execution, 0), minutes}
}

func (pq priorityQueue) Len() int { return len(pq.pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq.pq[i].upperGeodes(pq.bp, pq.minutes) > pq.pq[j].upperGeodes(pq.bp, pq.minutes)
}

func (pq priorityQueue) Swap(i, j int) { pq.pq[i], pq.pq[j] = pq.pq[j], pq.pq[i] }

func (pq *priorityQueue) Push(x *Execution) {
	pq.pq = append(pq.pq, x)
}

func (pq *priorityQueue) Pop() *Execution {
	old := pq.pq
	n := len(old)
	x := old[n-1]
	pq.pq = old[0 : n-1]
	return x
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func HeapPush(h *priorityQueue, x *Execution) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func HeapPop(h *priorityQueue) *Execution {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func up(h *priorityQueue, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down(h *priorityQueue, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

package day16

import "github.com/edwingeng/deque/v2"

type PartitionedQueue struct {
	queues               []*deque.Deque[State]
	partitions           int
	len                  int
	minQueueWithElements int
}

func NewPartitionedQueue(partitions int) *PartitionedQueue {
	queues := make([]*deque.Deque[State], partitions)
	for i := 0; i < partitions; i++ {
		queues[i] = deque.NewDeque[State]()
	}
	return &PartitionedQueue{
		queues:               queues,
		partitions:           partitions,
		minQueueWithElements: 1000,
	}
}

func (pq *PartitionedQueue) Push(x State) {
	i := pq.partitions - x.remainingMoves
	// HeapPush(&pq.queues[i], x)
	pq.queues[i].PushBack(x)
	if pq.minQueueWithElements > i {
		pq.minQueueWithElements = i
	}
	pq.len += 1
}

func (pq *PartitionedQueue) Pop() State {
	pq.len -= 1
	for i := pq.minQueueWithElements; i < pq.partitions; i++ {
		if pq.queues[i].Len() > 0 {
			// e := HeapPop(&pq.queues[i])
			e := pq.queues[i].PopFront()
			return e
		} else {
			pq.minQueueWithElements = i + 1
		}
	}
	panic("no elements in queue")
}

func (pq *PartitionedQueue) Len() int {
	return pq.len
}

type priorityQueue []State

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].Bound() > pq[j].Bound()
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x State) {
	*pq = append(*pq, x)
}

func (pq *priorityQueue) Pop() State {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// Push pushes the element x onto the heap.
// The complexity is O(log n) where n = h.Len().
func HeapPush(h *priorityQueue, x State) {
	h.Push(x)
	up(h, h.Len()-1)
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n) where n = h.Len().
// Pop is equivalent to Remove(h, 0).
func HeapPop(h *priorityQueue) State {
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

func createPriorityQueue() priorityQueue {
	return make(priorityQueue, 0, 1000)
}

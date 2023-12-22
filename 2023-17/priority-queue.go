package main

import "container/heap"

type PriorityQueue []*HeapBlock

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].minHeatLoss < pq[j].minHeatLoss {
		return true
	} else if pq[i].minHeatLoss == pq[j].minHeatLoss {
		if pq[i].address.y > pq[j].address.y {
			return true
		} else if pq[i].address.y == pq[j].address.y {
			if pq[i].address.x > pq[j].address.x {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	block := x.(*HeapBlock)
	contains, minHeatLoss := pq.contains(block)

	if contains {
		if minHeatLoss > block.minHeatLoss {
			pq.update(block, block.minHeatLoss)
		}
	} else {
		n := len(*pq)
		block.index = n
		*pq = append(*pq, block)
	}
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(block *HeapBlock, minHeatLoss int) {
	block.minHeatLoss = minHeatLoss
	heap.Fix(pq, block.index)
}

func (pq *PriorityQueue) contains(block *HeapBlock) (bool, int) {
	for _, b := range *pq {
		if b.address == block.address && int(b.directionToBlock) == int(block.directionToBlock) && b.stepsInDirection == block.stepsInDirection {
			return true, b.minHeatLoss
		}
	}
	return false, 0
}

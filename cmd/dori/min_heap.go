package main

type MinHeap struct {
	Length int
	data   []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		Length: 0,
		data:   []int{},
	}
}

func (heap *MinHeap) Insert(value int) {
	heap.data = append(heap.data, value)
	heap.heapifyUp(heap.Length)
	heap.Length++
}

func (heap *MinHeap) Delete() int {
	if heap.Length == 0 {
		return -1
	}

	heap.Length--
	out := heap.data[0]
	
	if heap.Length == 0 {
		heap.data = []int{}
		return out
	}

	heap.data[0] = heap.data[heap.Length]
	heap.heapifyDown(0)

	return out
}

func (heap *MinHeap) heapifyDown(idx int) {
	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if idx >= heap.Length || lIdx >= heap.Length {
		return
	}

	lV := heap.data[lIdx]
	rV := heap.data[rIdx]
	v := heap.data[idx]

	if lV > rV && v > rV {
		heap.data[idx] = rV
		heap.data[rIdx] = v
		heap.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		heap.data[idx] = lV
		heap.data[lIdx] = v
		heap.heapifyDown(lIdx)
	}
}

func (heap *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := parent(idx)
	parentV := heap.data[p]
	v := heap.data[idx]

	if parentV > v {
		heap.data[idx] = parentV
		heap.data[p] = v
		heap.heapifyUp(p)
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}

func leftChild(idx int) int {
	return idx*2 + 1
}

func rightChild(idx int) int {
	return idx*2 + 2
}

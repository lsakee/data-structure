package main

import "fmt"

type Heap[T comparable] struct {
	elements []T
	Compare  func(x, y T) bool
}

func (h *Heap[T]) Insert(value T) {
	if h.Length() == 0 {
		h.elements = make([]T, 0)
	}

	h.elements = append(h.elements, value)

	index := h.Length() - 1
	for !h.shiftUp(index) {
		switch index % 2 {
		case 0:
			index = (index - 2) / 2
		case 1:
			index = (index - 1) / 2
		}
	}
}

func (h *Heap[T]) Delete() (top T) {
	if h.Length() == 0 {
		return
	}

	root := h.elements[0]

	h.swap(0, h.Length()-1)
	h.elements = h.elements[:h.Length()-1]

	index := 0
	for next, done := h.shiftDown(index); !done; next, done = h.shiftDown(next) {
	}

	return root
}

func (h Heap[T]) Length() int {
	return len(h.elements)
}

func (h Heap[T]) Top() (top T) {
	top = h.elements[0]
	return
}

func (h Heap[T]) Print() {
	fmt.Printf("%v\n", h.elements)
}

func (h *Heap[T]) shiftUp(index int) (done bool) {
	if index == 0 {
		return true
	}

	current := h.elements[index]
	var parentIndex int

	if index%2 == 0 {
		parentIndex = (index - 2) / 2
	} else {
		parentIndex = (index - 1) / 2
	}

	parent := h.elements[parentIndex]

	if h.Compare(current, parent) {
		h.swap(index, parentIndex)
		return
	}

	return true
}

func (h *Heap[T]) shiftDown(index int) (nextIndex int, done bool) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	if leftChildIndex > h.Length()-1 {
		return -1, true
	} else if rightChildIndex > h.Length()-1 {
		if h.Compare(h.elements[leftChildIndex], h.elements[index]) {
			h.swap(leftChildIndex, index)
			return leftChildIndex, false
		}

		return -1, true
	}

	if h.Compare(h.elements[leftChildIndex], h.elements[rightChildIndex]) {
		if h.Compare(h.elements[leftChildIndex], h.elements[index]) {
			h.swap(leftChildIndex, index)
			return leftChildIndex, false
		}

		return -1, true
	} else {
		if h.Compare(h.elements[rightChildIndex], h.elements[index]) {
			h.swap(rightChildIndex, index)
			return rightChildIndex, false
		}
		return -1, true
	}
}

func (h *Heap[T]) swap(x, y int) {
	temp := h.elements[x]
	h.elements[x] = h.elements[y]
	h.elements[y] = temp
}

func main() {
	maxHeap := &Heap[int]{
		elements: make([]int, 0),
		Compare: func(x, y int) bool {
			return x > y
		},
	}

	maxHeap.Insert(7)
	maxHeap.Print() // [7]
	maxHeap.Insert(8)
	maxHeap.Print() // [8 7]
	maxHeap.Insert(1)
	maxHeap.Print() // [8 7 1]
	maxHeap.Insert(2)
	maxHeap.Print() // [8 7 1 2]
	maxHeap.Insert(3)
	maxHeap.Print() // [8 7 1 2 3]

	fmt.Println(maxHeap.Delete())
	maxHeap.Print()
	fmt.Println(maxHeap.Delete())
	maxHeap.Print()
}

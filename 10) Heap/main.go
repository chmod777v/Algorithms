package main

import "fmt"

type Heap struct {
	data []int
}

func NewHeap() *Heap {
	return &Heap{}
}

func (h *Heap) Push(x int) {
	h.data = append(h.data, x)
	i := len(h.data) - 1
	for i > 0 {
		p := (i - 1) / 2
		if h.data[p] <= h.data[i] {
			break

		}
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}

func (h *Heap) Pop() (int, bool) {
	if len(h.data) == 0 {
		return 0, false
	}
	top := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	n := len(h.data)
	h.siftDown(0, n)

	return top, true
}

func (h *Heap) siftDown(i, n int) {
	for {
		l, r := 2*i+1, 2*i+2
		s := i
		if l < n && h.data[l] < h.data[s] {
			s = l
		}
		if r < n && h.data[r] < h.data[s] {
			s = r
		}
		if s == i {
			break
		}
		h.data[i], h.data[s] = h.data[s], h.data[i]
		i = s
	}

}
func (h *Heap) Heapify(items []int) {
	h.data = append(h.data[:0], items...) // копия, чтобы не портить исходный срез
	n := len(h.data)
	for i := n/2 - 1; i >= 0; i-- {
		h.siftDown(i, n)
	}
}

func main() {
	heap := NewHeap()
	heap.Heapify([]int{11, 6, 10, 5, 3, 9, 8})
	fmt.Println(heap)
	heap.Push(4)
	fmt.Println(heap)

	if ans, ok := heap.Pop(); ok {
		fmt.Println(ans)
	}

	fmt.Println(heap)
}

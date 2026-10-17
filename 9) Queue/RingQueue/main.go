package main

import "fmt"

type RingQueue[T any] struct {
	buf   []T
	head  int
	tail  int
	count int
}

func NewRingQueue[T any](capacity int) *RingQueue[T] {
	return &RingQueue[T]{buf: make([]T, capacity)}
}

func (q *RingQueue[T]) Add(x T) {
	if q.count == len(q.buf) {
		q.grow()
	}
	q.buf[q.tail] = x
	q.tail = (q.tail + 1) % len(q.buf)
	q.count++
}

func (q *RingQueue[T]) Get() T {
	x := q.buf[q.head]
	q.head = (q.head + 1) % len(q.buf)
	q.count--
	return x
}

func (q *RingQueue[T]) grow() {
	newCap := len(q.buf) * 2
	if newCap == 0 {
		newCap = 1
	}
	newBuf := make([]T, newCap)
	for i := 0; i < q.count; i++ {
		newBuf[i] = q.buf[(q.head+i)%len(q.buf)]
	}
	q.buf = newBuf
	q.head = 0
	q.tail = q.count
}

func (q *RingQueue[T]) Len() int      { return q.count }
func (q *RingQueue[T]) IsEmpty() bool { return q.count == 0 }
func (q *RingQueue[T]) IsFull() bool  { return q.count == len(q.buf) }

func main() {
	q := NewRingQueue[int](3)
	q.Add(1)
	q.Add(2)
	q.Add(3)

	for !q.IsEmpty() {
		fmt.Printf("Длина: %d, Элемент:%v\n", q.Len(), q.Get())
	}
}

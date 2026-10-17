package main

import "fmt"

type Deque[T any] struct {
	buf   []T
	head  int // индекс первого элемента
	tail  int // индекс после последнего
	count int
}

func NewDeque[T any](capacity int) *Deque[T] {
	if capacity < 1 {
		capacity = 1
	}
	return &Deque[T]{buf: make([]T, capacity)}
}

// --- сдвиг индексов ---

func (q *Deque[T]) next(i int) int {
	return (i + 1) % len(q.buf)
}

func (q *Deque[T]) prev(i int) int {
	return (i - 1 + len(q.buf)) % len(q.buf)
}

// --- вставка ---

func (q *Deque[T]) PushBack(x T) {
	if q.count == len(q.buf) {
		q.grow()
	}
	q.buf[q.tail] = x
	q.tail = q.next(q.tail)
	q.count++
}

func (q *Deque[T]) PushFront(x T) {
	if q.count == len(q.buf) {
		q.grow()
	}
	q.head = q.prev(q.head)
	q.buf[q.head] = x
	q.count++
}

// --- удаление ---

func (q *Deque[T]) PopFront() T {
	x := q.buf[q.head]
	var zero T
	q.buf[q.head] = zero // для GC
	q.head = q.next(q.head)
	q.count--
	return x
}

func (q *Deque[T]) PopBack() T {
	q.tail = q.prev(q.tail)
	x := q.buf[q.tail]
	var zero T
	q.buf[q.tail] = zero // для GC
	q.count--
	return x
}

// --- доступ к концам ---

func (q *Deque[T]) Front() T {
	return q.buf[q.head]
}

func (q *Deque[T]) Back() T {
	return q.buf[q.prev(q.tail)]
}

// --- состояние ---

func (q *Deque[T]) Len() int      { return q.count }
func (q *Deque[T]) IsEmpty() bool { return q.count == 0 }
func (q *Deque[T]) IsFull() bool  { return q.count == len(q.buf) }

// --- расширение ---

func (q *Deque[T]) grow() {
	newCap := len(q.buf) * 2
	newBuf := make([]T, newCap)

	// Оставляем запас с обеих сторон, чтобы оба конца могли расти.
	shift := newCap/4 + 1

	n := copy(newBuf[shift:], q.buf[q.head:]) // от head до конца буфера
	copy(newBuf[shift+n:], q.buf[:q.tail])    // от начала до tail

	q.buf = newBuf
	q.head = shift
	q.tail = shift + q.count
}

// --- демонстрация ---

func main() {
	d := NewDeque[int](4)

	d.PushBack(1)
	d.PushBack(2)
	d.PushFront(0)
	d.PushFront(-1)

	fmt.Println("front:", d.Front()) // -1
	fmt.Println("back: ", d.Back())  // 2

	for !d.IsEmpty() {
		fmt.Println(d.PopFront())
	}

	d.PushBack(10)
	d.PushBack(20)
	d.PushBack(30)
	fmt.Println("pop back:", d.PopBack())   // 30
	fmt.Println("pop front:", d.PopFront()) // 10
}

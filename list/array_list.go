package list

import (
	"github.com/charlienet/go-datastructures/locker"
)

const minCapacity = 16

type arrayList[T any] struct {
	buf    []T
	minCap int
	head   int
	tail   int
	size   int
	locker.Locker
}

func NewArrayList[T any](values ...T) *arrayList[T] {
	minCap := minCapacity
	size := len(values)
	for minCap < size {
		minCap <<= 1
	}

	var tail int = size
	var buf []T
	if len(values) > 0 {
		buf := make([]T, minCap)
		copy(buf, values)
	}

	return &arrayList[T]{
		buf:  buf,
		tail: tail}
}

func (l *arrayList[T]) PushFront(value T) {
	l.Lock()
	defer l.Unlock()

	l.grow()

	l.head = l.prev(l.head)
	l.buf[l.head] = value
	l.size++
}

func (l *arrayList[T]) PushBack(v T) {
	l.Lock()
	defer l.Unlock()

	l.grow()

	l.buf[l.tail] = v

	l.tail = l.next(l.tail)
	l.size++
}

func (l *arrayList[T]) PopFront() T {
	l.Lock()
	defer l.Unlock()

	if l.size <= 0 {
		panic("list: PopFront() called on empty list")
	}
	ret := l.buf[l.head]
	var zero T
	l.buf[l.head] = zero

	l.head = l.next(l.head)
	l.size--

	l.shrink()
	return ret
}

func (l *arrayList[T]) PopBack() T {
	l.Lock()
	defer l.Unlock()

	l.tail = l.prev(l.tail)

	ret := l.buf[l.tail]
	var zero T
	l.buf[l.tail] = zero
	l.size--

	l.shrink()
	return ret
}

func (l *arrayList[T]) RemoveAt(at int) T {
	if at < 0 || at >= l.size {
		panic(ErrorOutOffRange)
	}

	l.Lock()
	defer l.Unlock()

	rm := (l.head + at) & (len(l.buf) - 1)
	if at*2 < l.size {
		for i := 0; i < at; i++ {
			prev := l.prev(rm)
			l.buf[prev], l.buf[rm] = l.buf[rm], l.buf[prev]
			rm = prev
		}
		return l.PopFront()
	}
	swaps := l.size - at - 1
	for i := 0; i < swaps; i++ {
		next := l.next(rm)
		l.buf[rm], l.buf[next] = l.buf[next], l.buf[rm]
		rm = next
	}
	return l.PopBack()
}

func (l *arrayList[T]) Front() T {
	l.Lock()
	defer l.Unlock()

	return l.buf[l.head]
}

func (l *arrayList[T]) Back() T {
	l.Lock()
	defer l.Unlock()

	return l.buf[l.tail]
}

func (l *arrayList[T]) ForEach(fn func(T)) {
	l.Lock()
	defer l.Unlock()

	n := l.head
	for i := 0; i < l.size; i++ {
		fn(l.buf[n])

		n = l.next(n)
	}
}

func (q *arrayList[T]) prev(i int) int {
	return (i - 1) & (len(q.buf) - 1)
}

func (l *arrayList[T]) next(i int) int {
	return (i + 1) & (len(l.buf) - 1)
}

func (l *arrayList[T]) grow() {
	if l.size != len(l.buf) {
		return
	}
	if len(l.buf) == 0 {
		if l.minCap == 0 {
			l.minCap = minCapacity
		}
		l.buf = make([]T, l.minCap)
		return
	}

	l.resize()
}

func (l *arrayList[T]) shrink() {
	if len(l.buf) > l.minCap && (l.size<<2) == len(l.buf) {
		l.resize()
	}
}

// resize resizes the list to fit exactly twice its current contents. This is
// used to grow the list when it is full, and also to shrink it when it is
// only a quarter full.
func (l *arrayList[T]) resize() {
	newBuf := make([]T, l.size<<1)
	if l.tail > l.head {
		copy(newBuf, l.buf[l.head:l.tail])
	} else {
		n := copy(newBuf, l.buf[l.head:])
		copy(newBuf[n:], l.buf[:l.tail])
	}

	l.head = 0
	l.tail = l.size
	l.buf = newBuf
}

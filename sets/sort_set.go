package sets

import (
	"cmp"
	"iter"
	"slices"

	"github.com/charlienet/go-datastructures/locker"
)

type sorted_set[T cmp.Ordered] struct {
	sorted []T
	m      *hash_set[T]
	locker.Locker
}

func NewSortedSet[T cmp.Ordered](ss ...T) *sorted_set[T] {
	return &sorted_set[T]{
		m: New(ss...),
	}
}

func (s *sorted_set[T]) Add(values ...T) *sorted_set[T] {
	for _, v := range values {
		s.m.Add(v)
	}

	return s
}

func (s *sorted_set[T]) Remove(values ...T) *sorted_set[T] {
	for _, v := range values {
		s.m.Remove(v)
	}

	return s
}

func (s *sorted_set[T]) Exists(i T) bool {
	return s.m.Exists(i)
}

func (s *sorted_set[T]) Clear() *sorted_set[T] {
	s.m.Clear()
	s.sorted = nil

	return s
}

func (s *sorted_set[T]) Len() int {
	return s.m.Size()
}

func (s *sorted_set[T]) Asc() *sorted_set[T] {

	s.sorted = s.m.Keys()
	slices.Sort(s.sorted)

	return s
}

func (s *sorted_set[T]) Desc() *sorted_set[T] {
	slices.SortFunc(s.sorted, func(a, b T) int {
		if a == b {
			return 0
		}

		if a > b {
			return 1
		}

		return -1
	})

	return s
}

func (s *sorted_set[T]) Iterator() iter.Seq[T] {

	return func(yield func(T) bool) {
		for _, v := range s.sorted {
			if !yield(v) {
				break
			}
		}
	}
}

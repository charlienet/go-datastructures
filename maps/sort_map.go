package maps

import (
	"cmp"
	"iter"
	"slices"

	smaps "maps"

	"github.com/charlienet/go-datastructures/locker"
)

type SortedMap[K cmp.Ordered, V any] struct {
	sorted []K
	m      map[K]V
	locker.Locker
}

func NewSortedMap[K cmp.Ordered, V any](maps ...map[K]V) *SortedMap[K, V] {
	merged := Merge(maps...)

	keys := make([]K, 0, len(merged))

	return &SortedMap[K, V]{
		sorted: slices.AppendSeq(keys, smaps.Keys(merged)),
		m:      merged,
	}
}

func (s *SortedMap[K, V]) Set(k K, v V) {
	s.Lock()
	defer s.Unlock()

	s.m[k] = v

	keys := make([]K, 0, len(s.m))
	s.sorted = slices.AppendSeq(keys, smaps.Keys(s.m))
}

func (s *SortedMap[K, V]) Get(k K) (V, bool) {
	s.Lock()
	defer s.Unlock()

	v, ok := s.m[k]
	return v, ok
}

func (s *SortedMap[K, V]) Delete(k K) {
	s.Lock()
	defer s.Unlock()

	delete(s.m, k)
}

func (s *SortedMap[K, V]) Asc() *SortedMap[K, V] {
	s.Lock()
	defer s.Unlock()

	s.sortKeys()
	slices.Sort(s.sorted)
	return s
}

func (s *SortedMap[K, V]) Desc() *SortedMap[K, V] {
	s.Lock()
	defer s.Unlock()

	s.sortKeys()
	slices.SortFunc(s.sorted, func(a, b K) int {
		if a == b {
			return 0
		}
		if a < b {
			return 1
		}
		return -1
	})

	return s
}

func (s *SortedMap[K, V]) All() iter.Seq2[K, V] {
	s.Lock()
	defer s.Unlock()

	return func(yield func(K, V) bool) {
		for _, k := range s.sorted {
			if !yield(k, s.m[k]) {
				break
			}
		}
	}
}

func (s *SortedMap[K, V]) sortKeys() {
	keys := make([]K, 0, len(s.m))
	s.sorted = slices.AppendSeq(keys, smaps.Keys(s.m))
}

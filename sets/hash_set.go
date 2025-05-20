package sets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/charlienet/go-datastructures/locker"
)

type hash_set[T comparable] struct {
	m map[T]struct{}
	l locker.Locker
}

func New[T comparable](values ...T) *hash_set[T] {
	set := &hash_set[T]{
		m: make(map[T]struct{}, len(values)),
	}

	set.Add(values...)
	return set
}

func (s *hash_set[T]) Synchronize() *hash_set[T] {
	s.l.Synchronize()

	return s
}

func (s *hash_set[T]) Add(values ...T) *hash_set[T] {
	s.l.Lock()

	for _, v := range values {
		s.m[v] = struct{}{}
	}

	s.l.Unlock()

	return s
}

func (s *hash_set[T]) Remove(values ...T) *hash_set[T] {
	s.l.Lock()

	for _, v := range values {
		delete(s.m, v)
	}

	s.l.Unlock()

	return s
}

func (s *hash_set[T]) Keys() []T {
	s.l.RLock()
	defer s.l.RUnlock()

	keys := make([]T, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}

	return keys
}

func (s *hash_set[T]) Exists(i T) bool {
	s.l.RLock()
	defer s.l.RUnlock()

	_, ok := s.m[i]
	return ok
}

func (s *hash_set[T]) Clear() *hash_set[T] {
	s.l.Lock()

	clear(s.m)

	s.l.Unlock()
	return s
}

func (s *hash_set[T]) Size() int {
	return len(s.m)
}

func (s *hash_set[T]) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Size())

	for ele := range s.m {
		b, err := json.Marshal(ele)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ", "))), nil
}

func (s *hash_set[T]) UnmarshalJSON(b []byte) error {
	var i []any

	d := json.NewDecoder(bytes.NewReader(b))
	d.UseNumber()
	err := d.Decode(&i)
	if err != nil {
		return err
	}

	for _, v := range i {
		if t, ok := v.(T); ok {
			s.Add(t)
		}
	}

	return nil
}

func (s *hash_set[T]) String() string {
	l := make([]string, 0, len(s.m))
	for k := range s.m {
		l = append(l, fmt.Sprint(k))
	}

	return fmt.Sprintf("{%s}", strings.Join(l, ", "))
}

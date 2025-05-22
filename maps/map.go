package maps

import (
	"cmp"
	"iter"
	smaps "maps"
	"strings"
)

type Map[K comparable, V any] interface {
	Iter() iter.Seq2[K, V]
	Count() int
}

func Merge[K comparable, V any](maps ...map[K]V) map[K]V {

	totalCount := 0
	for _, m := range maps {
		totalCount += len(m)
	}

	ret := make(map[K]V, totalCount)
	for _, m := range maps {
		smaps.Copy(ret, m)
	}

	return ret
}

func Asc[K cmp.Ordered, V any](m map[K]V) Map[K, V] {
	return NewSortedMap(m).Asc()
}

func Desc[K cmp.Ordered, V any](m map[K]V) Map[K, V] {
	return NewSortedMap(m).Desc()
}

func Join[K comparable, V any](m Map[K, V], sep string, f func(k K, v V) string) string {
	slice := make([]string, 0, m.Count())

	for k, v := range m.Iter() {
		slice = append(slice, f(k, v))
	}

	return strings.Join(slice, sep)
}

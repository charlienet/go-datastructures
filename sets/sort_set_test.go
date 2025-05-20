package sets_test

import (
	"testing"

	"github.com/charlienet/go-datastructures/sets"
)

func TestSort(t *testing.T) {
	s := sets.NewSortedSet("a", "c", "e", "b", "d")

	s.Synchronize()

	s.Add("f", "h", "g")
	t.Logf("%v", s)

	for n := range s.Asc().Iterator() {
		t.Logf("%v", n)
	}
}

func BenchmarkSort(b *testing.B) {

	for b.Loop() {
		s := sets.NewSortedSet("a", "c", "e", "b", "d")
		s.Add("f", "h", "g")
	}
}

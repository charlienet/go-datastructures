package maps

import "testing"

func TestSort(t *testing.T) {
	m := NewSortedMap(map[string]string{"a": "a", "c": "c", "e": "e", "b": "b", "d": "d"})

	m.Synchronize()

	for n, v := range m.Asc().All() {
		t.Logf("%v %v", n, v)
	}

	for n, v := range m.Desc().All() {
		t.Logf("%v %v", n, v)
	}
}

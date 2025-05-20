package sets_test

import (
	"testing"

	"github.com/charlienet/go-datastructures/sets"
)

func TestUnique(t *testing.T) {
	r := sets.Unique("a", "b", "a", "c", "b", "e")
	t.Logf("%v", r)
}



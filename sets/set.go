package sets

type Set[T comparable] interface {
}

func Unique[T comparable](values ...T) []T {
	r := make([]T, 0, len(values))
	s := New(values...)

	for _, v := range values {
		if s.Exists(v) {
			s.Remove(v)
			r = append(r, v)
		}
	}

	return r
}

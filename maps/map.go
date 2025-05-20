package maps

import smaps "maps"

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

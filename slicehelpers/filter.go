package slicehelpers

func FilterSlice[T comparable](s []T, filterFunc func(T) bool) []T {
	result := []T{}
	for _, v := range s {
		if filterFunc(v) {
			result = append(result, v)
		}
	}
	return result
}

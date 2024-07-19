package slicehelpers

func Remove[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func RemoveIndices[T any](s []T, indices ...int) []T {
	for _, i := range indices {
		s = Remove(s, i)
	}
	return s
}

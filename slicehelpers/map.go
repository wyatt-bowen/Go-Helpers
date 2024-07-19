package slicehelpers

func Map[T any](s []T, mapFunction func(T) T) []T {
	result := []T{}
	for _, v := range s {
		result = append(result, mapFunction(v))
	}
	return result
}

func MapInplace[T any](s []T, mapFunction func(T) T) {
	for i, v := range s {
		s[i] = mapFunction(v)
	}
}

func Square() func(fromSlice int) int {
	return func(fromSlice int) int { return fromSlice * fromSlice }
}

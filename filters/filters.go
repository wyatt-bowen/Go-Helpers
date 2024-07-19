package filters

// Typical Filters

func Equals[T comparable](other T) func(fromSlice T) bool {
	return func(fromSlice T) bool { return fromSlice == other }
}

func BetweenInclusive(low, high int) func(fromSlice int) bool {
	return func(fromSlice int) bool { return fromSlice >= low && fromSlice <= high }
}

func Even() func(fromSlice int) bool {
	return func(fromSlice int) bool { return fromSlice%2 == 0 }
}

func Or[T comparable](predicates ...func(fromSlice T) bool) func(T) bool {
	return func(fromSlice T) bool {
		result := false
		for _, predicate := range predicates {
			result = result || predicate(fromSlice)
		}
		return result
	}
}

func Xor[T comparable](predicates ...func(fromSlice T) bool) func(T) bool {
	return func(fromSlice T) bool {
		result := false
		for _, predicate := range predicates {
			result = result != predicate(fromSlice)
		}
		return result
	}
}

func And[T comparable](predicates ...func(fromSlice T) bool) func(T) bool {
	return func(fromSlice T) bool {
		result := true
		for _, predicate := range predicates {
			result = result && predicate(fromSlice)
		}
		return result
	}
}

func Not[T comparable](predicate func(T) bool) func(T) bool {
	return func(fromSlice T) bool { return !predicate(fromSlice) }
}

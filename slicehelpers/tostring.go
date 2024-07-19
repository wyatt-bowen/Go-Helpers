package slicehelpers

import "fmt"

func ToString[T any](s []T) string {
	result := ""
	for i := 0; i < len(s)-1; i++ {
		result += fmt.Sprintf("%v, ", s[i])
	}
	result += fmt.Sprintf("%v", s[len(s)-1])

	return result
}

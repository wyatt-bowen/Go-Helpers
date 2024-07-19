package stack

import "fmt"

type Stack[T any] struct {
	data   []T
	length int
}

func (s *Stack[T]) IsEmpty() bool {
	return s.length == 0
}

func (s *Stack[T]) Push(values ...T) {
	s.length++
	s.data = append(s.data, values...)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	} else {
		result := s.data[s.length-1]
		s.data = s.data[:s.length-1]
		s.length--
		return result, true
	}
}

func (s *Stack[T]) Top() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	} else {
		result := s.data[s.length-1]
		return result, true
	}
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0), 0}
}

func NewStackFrom[T any](s ...T) *Stack[T] {
	result := NewStack[T]()
	for _, v := range s {
		result.Push(v)
	}
	return result
}

func (s *Stack[T]) Length() int {
	return s.length
}

func (s *Stack[T]) ToString() string {
	result := ""
	for i := 0; i < len(s.data)-1; i++ {
		result += fmt.Sprintf("%v, ", s.data[i])
	}
	result += fmt.Sprintf("%v", s.data[s.length-1])

	return result
}

func (s *Stack[T]) GetData() *[]T {
	return &s.data
}

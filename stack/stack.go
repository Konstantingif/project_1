
package stack


type Stack[T any] struct {
	items []T
}


func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}


func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item, true
}


func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}


func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}


func (s *Stack[T]) Len() int {
	return len(s.items)
}
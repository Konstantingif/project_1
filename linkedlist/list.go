package linkedlist

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	size int
}

func (l *List[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	l.size++
}

func (l *List[T]) Prepend(value T) {
	newNode := &Node[T]{Value: value, Next: l.Head}
	l.Head = newNode
	l.size++
}

func (l *List[T]) IsEmpty() bool {
	return l.Head == nil
}

func (l *List[T]) Len() int {
	return l.size
}

func (l *List[T]) Values() []T {
	var result []T
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

package validparentheses

type stack[T comparable] []T

func (s *stack[T]) push(r T) {
	*s = append(*s, r)
}

func (s *stack[T]) pop() (h T) {
	h = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

package validparentheses

type stack []rune

func (s *stack) push(r rune) {
	*s = append(*s, r)
}

func (s *stack) pop() (h rune) {
	h = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}

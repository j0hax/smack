package lexer

type Stack []Token

// Empty returns true if the length of the stack is zero.
func (s *Stack) Empty() bool {
	return len(*s) == 0
}

// Push adds a value to the stack
func (s *Stack) Push(n Token) {
	*s = append(*s, n)
}

// Pop removes the top token from the stack and returns it
func (s *Stack) Pop() *Token {
	index := len(*s) - 1
	element := (*s)[index]
	// Chop off the last part
	*s = (*s)[:index]
	return &element
}

// Peek returns the top element of the stack without removing it
func (s *Stack) Peek() *Token {
	index := len(*s) - 1
	return &(*s)[index]
}

// PopN pops n items from the stack
func (s *Stack) PopN(count int) Stack {
	var newStack Stack
	for i := 0; i < count; i++ {
		newStack.Push(*s.Pop())
	}

	return newStack
}

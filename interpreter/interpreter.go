package interpreter

import (
	"fmt"
	"strings"

	"github.com/j0hax/smack/lexer"
)

var s lexer.Stack

// Process lexes as string and processes the tokens on the stack,
// returning the topmost token
func Process(input string) *lexer.Token {
	t := lexer.LexString(input)

	for _, o := range t {
		if o.Type == lexer.Function {
			tokens := s.PopN(o.NumArgs)
			result := o.Compute(tokens...)
			s.Push(result)
		} else if o.Type == lexer.Number {
			s.Push(o)
		}
	}

	return s.Peek()
}

// Stacktrace returns a representation of the current Stack in descending order.
func StackTrace() string {
	var items []string
	l := len(s) - 1

	for i := l; i >= 0; i-- {
		t := s[i]
		item := fmt.Sprintf("%X\t%s", i, t.String())
		items = append(items, item)
	}

	return strings.Join(items, "\n")
}

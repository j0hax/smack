package lexer

import (
	"fmt"
)

// TokenType is an enumeration for different kinds of tokens
type TokenType int

const (
	Number TokenType = iota
	Function
)

// TokenFun is a function that can process multiple other tokens as arguments.
//
// This function is called for different operators (addition, subtraction, etc.).
type TokenFun func(tokens ...Token) Token

// A token is a lexed input.
//
// It can have multiple meanings depending on the type.
type Token struct {
	Type    TokenType
	Data    any
	NumArgs int
	Compute TokenFun
}

// NewNumberToken creates a new token representing a number
func NewNumberToken(n float64) *Token {
	return &Token{
		Type: Number,
		Data: n,
	}
}

func (t *Token) String() string {
	switch t.Type {
	case Number:
		return fmt.Sprint(t.Data.(float64))
	case Function:
		return t.Data.(string)
	default:
		return fmt.Sprint(t.Data)
	}
}

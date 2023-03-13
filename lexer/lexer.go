package lexer

import (
	"strconv"
	"strings"
)

// LexSymbol interprets an input symbol as a single token
func LexSymbol(s string) *Token {
	switch s {
	case "+":
		return AdditionToken
	case "-":
		return SubtractionToken
	case "*":
		return MultiplicationToken
	case "/":
		return DivisionToken

	case "pow":
		return PowToken

	case "log":
		return LogToken

	case "inv":
		return InvToken

	// For now, assume the input is always a number (unsafe)
	default:
		n, _ := strconv.ParseFloat(s, 64)
		return &Token{
			Type:    Number,
			Data:    n,
			NumArgs: 0,
		}
	}
}

// LexString interprets an input program and returns a ready-to-use stack
func LexString(input string) Stack {
	in := strings.Split(input, " ")

	var stack Stack
	for _, symbol := range in {
		t := LexSymbol(symbol)
		stack.Push(*t)
	}

	return stack
}

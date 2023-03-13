package lexer

import "math"

// LogToken consumes two tokens, first a base, then a number
var LogToken = &Token{
	Type:    Function,
	NumArgs: 2,
	Compute: func(tokens ...Token) Token {
		base := tokens[1].Data.(float64)
		data := tokens[0].Data.(float64)
		result := math.Log(data) / math.Log(base)
		return *NewNumberToken(result)
	}}

// PowToken consumes two tokens, first a base, its exponent
var PowToken = &Token{
	Type:    Function,
	NumArgs: 2,
	Compute: func(tokens ...Token) Token {
		base := tokens[1].Data.(float64)
		data := tokens[0].Data.(float64)
		result := math.Pow(base, data)
		return *NewNumberToken(result)
	}}

// InvToken negates a number
var InvToken = &Token{
	Type:    Function,
	NumArgs: 1,
	Compute: func(tokens ...Token) Token {
		base := tokens[0].Data.(float64)
		return *NewNumberToken(-base)
	}}

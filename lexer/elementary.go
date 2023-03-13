package lexer

var AdditionToken = &Token{
	Type:    Function,
	NumArgs: 2,
	Compute: func(tokens ...Token) Token {
		intermediate := float64(0)
		for _, t := range tokens {
			intermediate = t.Data.(float64) + intermediate
		}
		return *NewNumberToken(intermediate)
	}}

var SubtractionToken = &Token{
	Type:    Function,
	NumArgs: 2,
	Compute: func(tokens ...Token) Token {
		intermediate := float64(0)
		for _, t := range tokens {
			intermediate = t.Data.(float64) - intermediate
		}
		return *NewNumberToken(intermediate)
	}}

var MultiplicationToken = &Token{
	Type:    Function,
	NumArgs: 2,
	Compute: func(tokens ...Token) Token {
		intermediate := float64(1)
		for _, t := range tokens {
			intermediate = t.Data.(float64) * intermediate
		}
		return *NewNumberToken(intermediate)
	}}

var DivisionToken = &Token{
	Type:    Function,
	NumArgs: 2,
	Compute: func(tokens ...Token) Token {
		intermediate := float64(1)
		for _, t := range tokens {
			intermediate = t.Data.(float64) / intermediate
		}
		return *NewNumberToken(intermediate)
	}}

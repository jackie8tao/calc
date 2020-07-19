package main

type Token int

const (
	TokInt    Token = iota + 1 // integer
	TokAdd                     // '+'
	TokSub                     // '-'
	TokMul                     // '*'
	TokDiv                     // '/'
	TokLParen                  // '('
	TokRParen                  // ')'
)

// lookup the token upon char
func lookupToken(ch rune) Token {
	var tok Token
	switch string(ch) {
	case "+":
		tok = TokAdd
	case "-":
		tok = TokSub
	case "*":
		tok = TokMul
	case "/":
		tok = TokDiv
	case "(":
		tok = TokLParen
	case ")":
		tok = TokRParen
	default:
		panic(ErrIllegalToken)
	}
	return tok
}

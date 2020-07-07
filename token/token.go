package token

type Token int

const (
	ILLEGAL Token = iota
	INT           // integer
	ADD           // '+'
	SUB           // '-'
	MUL           // '*'
	DIV           // '/'
	LPAREN        // '('
	RPAREN        // ')'
)

func LookupOperator(ch rune) Token {
	var tok Token
	switch string(ch) {
	case "+":
		tok = ADD
	case "-":
		tok = SUB
	case "*":
		tok = MUL
	case "/":
		tok = DIV
	default:
		tok = ILLEGAL
	}
	return tok
}

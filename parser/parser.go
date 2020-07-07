package parser

import "calc/lexer"

type Parser struct {
	err   error
	lexer lexer.Lexer
}

func New() *Parser {
	return &Parser{}
}

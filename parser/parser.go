package parser

import (
	"errors"
	"strconv"

	"calc/ast"
	"calc/lexer"
	"calc/token"
)

var (
	errUnexpectedToken = errors.New("unexpected token")
)

type Parser struct {
	lexer lexer.Lexer
	tok   token.Token
	val   string
	err   error
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) eat() {
	tok, val := p.lexer.Get()
	if p.lexer.Err() != nil {
		panic(p.lexer.Err())
	}
	p.tok = tok
	p.val = val
}

func (p *Parser) match(tok token.Token) {
	p.eat()
	if p.tok != tok {
		panic(errUnexpectedToken)
	}
}

func (p *Parser) prefixTerm() (ret ast.Ast) {
	switch p.tok {
	case token.MUL:
		p.eat()
		factor := p.factor()
	}
}

func (p *Parser) factor() (ret ast.Ast) {
	switch p.tok {
	case token.INT:
		val, err := strconv.ParseInt(p.val, 10, 64)
		if err != nil {
			panic(err)
		}
		ret = ast.NewIntExpr(val)
	case token.LPAREN:
		p.eat()
		ret = p.expr()
		p.match(token.RPAREN)
	default:
		panic(errUnexpectedToken)
	}
	return
}

func (p *Parser) prefixExpr() ast.Ast {

}

func (p *Parser) term() ast.Ast {

}

func (p *Parser) expr() ast.Ast {

}

func (p *Parser) Next() ast.Ast {

}

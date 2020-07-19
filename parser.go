package main

import (
	"strconv"
)

// Parser calculator parser.
type Parser struct {
	err error
	lx  *Lexer
	tok Token
	val string
}

// NewParser create parser.
func NewParser(src string) *Parser {
	lx := NewLexer(src)
	tk, val := lx.Get()
	if lx.Err() != nil {
		panic(lx.Err())
	}
	return &Parser{
		lx:  lx,
		tok: tk,
		val: val,
	}
}

// get next token
func (p *Parser) eat() {
	tok, val := p.lx.Get()
	if p.lx.Err() != nil {
		if p.lx.Err() != ErrEOZ {
			panic(p.lx.Err())
		}
		// when meets EOZ, we simply set tok and val zero.
		p.tok = Token(0)
		p.val = "$"
		return
	}
	p.tok = tok
	p.val = val
}

// get next token and check whether it is tok.
func (p *Parser) match(tok Token) {
	p.eat()
	if p.tok != tok {
		panic(ErrUnexpectedToken)
	}
}

// parse prefix-term
func (p *Parser) prefixTerm(inh Ast) (ret Ast) {
	var factor Ast
	switch p.tok {
	case TokMul:
		p.eat()
		factor = p.factor()
		ret = &MulAst{
			Left:  inh,
			Right: factor,
		}
	case TokDiv:
		p.eat()
		factor = p.factor()
		ret = &DivAst{
			Left:  inh,
			Right: factor,
		}
	default: // return nil
		return
	}

	pfx := p.prefixTerm(ret)
	if pfx != nil {
		ret = pfx
	}
	return
}

// parse factor.
func (p *Parser) factor() (ast Ast) {
	switch p.tok {
	case TokInt:
		val, err := strconv.ParseInt(p.val, 10, 64)
		if err != nil {
			panic(err)
		}
		ast = &IntAst{Num: val}
		p.eat()
	case TokLParen:
		p.eat()
		ast = p.expr()
		p.match(TokRParen)
	default:
		panic(ErrUnexpectedToken)
	}
	return
}

// parse term
func (p *Parser) term() (ret Ast) {
	factor := p.factor()
	ret = p.prefixTerm(factor)
	if ret != nil {
		return
	}
	ret = factor
	return
}

// parse prefix-term
func (p *Parser) prefixExpr(inh Ast) (ret Ast) {
	var term Ast
	switch p.tok {
	case TokAdd:
		p.eat()
		term = p.term()
		ret = &AddAst{
			Left:  inh,
			Right: term,
		}
	case TokSub:
		p.eat()
		term := p.term()
		ret = &SubAst{
			Left:  inh,
			Right: term,
		}
	default:
		return
	}

	pfx := p.prefixExpr(ret)
	if pfx != nil {
		ret = pfx
	}
	return
}

// parse expresion
func (p *Parser) expr() (ret Ast) {
	term := p.term()
	ret = p.prefixExpr(term)
	if ret != nil {
		return
	}
	ret = term
	return
}

// Next parse the ast from input tokens
func (p *Parser) Next() Ast {
	return p.expr()
}

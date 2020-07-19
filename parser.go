package main

import (
	"go/token"
	"strconv"

	"calc/ast"
)

// Node parsing tree node
type Node struct {
	Val ast.Ast
	Ins ast.Ast
}

type Parser struct {
	err error
	lx  *Lexer
	tok Token
	val string
}

// New create
func New(src string) *Parser {
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

// TODO we should set token value when meets the ErrEOZ
func (p *Parser) eat() {
	tok, val := p.lx.Get()
	if p.lx.Err() != nil {
		if p.lx.Err() != lexer.ErrEOZ {
			panic(p.lx.Err())
		}
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

func (p *Parser) prefixTerm(n *Node) ast.Ast {
	switch p.tok {
	case token.MUL:
		p.eat()
		factor := p.factor(nil)
		tmpAst := ast.NewMulExpr(n.Ins, factor)
		pfxTermAst := p.prefixTerm(&Node{Ins: tmpAst})
		if pfxTermAst != nil {
			return pfxTermAst
		}
		return tmpAst
	case token.DIV:
		p.eat()
		factor := p.factor(nil)
		tmpAst := ast.NewDivExpr(n.Ins, factor)
		pfxTermAst := p.prefixTerm(&Node{Ins: tmpAst})
		if pfxTermAst != nil {
			return pfxTermAst
		}
		return tmpAst
	}
	return nil
}

func (p *Parser) factor(_ *Node) (ret ast.Ast) {
	switch p.tok {
	case token.INT:
		val, err := strconv.ParseInt(p.val, 10, 64)
		if err != nil {
			panic(err)
		}
		ret = ast.NewIntExpr(val)
		p.eat()
	case token.LPAREN:
		p.eat()
		ret = p.expr()
		p.match(token.RPAREN)
	default:
		panic(errUnexpectedToken)
	}
	return
}

func (p *Parser) prefixExpr(n *Node) ast.Ast {
	switch p.tok {
	case token.ADD:
		p.eat()
		term := p.term(nil)
		tmpAst := ast.NewAddExpr(n.Ins, term)
		pfxExprAst := p.prefixExpr(&Node{Ins: tmpAst})
		if pfxExprAst != nil {
			return pfxExprAst
		}
		return tmpAst
	case token.SUB:
		p.eat()
		term := p.term(nil)
		tmpAst := ast.NewSubExpr(n.Ins, term)
		pfxExprAst := p.prefixExpr(&Node{Ins: tmpAst})
		if pfxExprAst != nil {
			return pfxExprAst
		}
	}
	return nil
}

func (p *Parser) term(_ *Node) ast.Ast {
	factor := p.factor(&Node{})
	pfxTermNode := &Node{Ins: factor}
	pfxTermAst := p.prefixTerm(pfxTermNode)
	if pfxTermAst != nil {
		return pfxTermAst
	}
	return factor
}

func (p *Parser) expr() ast.Ast {
	termNode, pfxExpNode := &Node{}, &Node{}
	termAst := p.term(termNode)
	pfxExpNode.Ins = termAst
	pfxExprAst := p.prefixExpr(pfxExpNode)
	if pfxExprAst != nil {
		return pfxExprAst
	}
	return termAst
}

// Next parse the ast from input tokens
func (p *Parser) Next() ast.Ast {
	return p.expr()
}

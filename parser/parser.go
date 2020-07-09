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

type Node struct {
	Val ast.Ast
	Ins ast.Ast
}

type Parser struct {
	err error
	lx  lexer.Lexer
	tok token.Token
	val string
}

func New() *Parser {
	return &Parser{}
}

func (p *Parser) eat() {
	tok, val := p.lx.Get()
	if p.lx.Err() != nil {
		panic(p.lx.Err())
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

func (p *Parser) prefixTerm(n *Node) (ret ast.Ast) {
	switch p.tok {
	case token.MUL:
		p.eat()
		factor := p.factor(nil)
		tmpAst := ast.NewMulExpr(n.Ins, factor)
		ret = p.prefixTerm(&Node{Ins: tmpAst})
	case token.DIV:
		p.eat()
		factor := p.factor(nil)
		tmpAst := ast.NewDivExpr(n.Ins, factor)
		ret = p.prefixTerm(&Node{Ins: tmpAst})
	}
	return
}

func (p *Parser) factor(_ *Node) (ret ast.Ast) {
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

func (p *Parser) prefixExpr(n *Node) ast.Ast {
	switch p.tok {
	case token.ADD:
		p.eat()
		term := p.term(nil)
		tmpAst := ast.NewAddExpr(n.Ins, term)
		return p.prefixExpr(&Node{Ins: tmpAst})
	case token.SUB:
		p.eat()
		term := p.term(nil)
		tmpAst := ast.NewSubExpr(n.Ins, term)
		return p.prefixExpr(&Node{Ins: tmpAst})
	}
	return nil
}

func (p *Parser) term(_ *Node) ast.Ast {
	factor := p.factor(&Node{})
	pfxTermNode := &Node{Ins: factor}
	return p.prefixTerm(pfxTermNode)
}

func (p *Parser) expr() ast.Ast {
	termNode, pfxExpNode := &Node{}, &Node{}
	termAst := p.term(termNode)
	pfxExpNode.Ins = termAst
	return p.prefixExpr(pfxExpNode)
}

func (p *Parser) Next() ast.Ast {

}

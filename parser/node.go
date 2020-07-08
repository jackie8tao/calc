package parser

import "calc/ast"

type Node interface {
	Ins() ast.Ast
	Val() ast.Ast
}

type (
	// TermNode
	TermNode struct {
		val ast.Ast
	}

	// FactorNode
	FactorNode struct {
		val ast.Ast
	}

	// PrefixTermNode
	PrefixTermNode struct {
		val ast.Ast
		ins ast.Ast
	}
)

func (n *TermNode) Ins() ast.Ast {
	return nil
}

func (n *TermNode) Val() ast.Ast {
	return n.val
}

func (n *FactorNode) Ins() ast.Ast {
	return nil
}

func (n *FactorNode) Val() ast.Ast {
	return n.val
}

func (n *PrefixTermNode) Ins() ast.Ast {
	return n.ins
}

func (n *PrefixTermNode) Val() ast.Ast {
	return n.val
}

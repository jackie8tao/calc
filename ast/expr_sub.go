package ast

type SubExpr struct {
	Left  Ast
	Right Ast
}

func NewSubExpr(left, right Ast) *SubExpr {
	return &SubExpr{
		Left:  left,
		Right: right,
	}
}

func (e *SubExpr) Walk() int64 {
	return e.Left.Walk() - e.Right.Walk()
}

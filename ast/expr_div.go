package ast

type DivExpr struct {
	Left  Ast
	Right Ast
}

func NewDivExpr(left, right Ast) *DivExpr {
	return &DivExpr{
		Left:  left,
		Right: right,
	}
}

func (e *DivExpr) Walk() int64 {
	factor := e.Right.Walk()
	if factor == 0 {
		panic("division by zero")
	}
	return e.Left.Walk() / factor
}

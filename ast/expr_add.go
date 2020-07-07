package ast

type AddExpr struct {
	Left  Ast
	Right Ast
}

func NewAddExpr(left, right Ast) *AddExpr {
	return &AddExpr{
		Left:  left,
		Right: right,
	}
}

func (e *AddExpr) Walk() int64 {
	return e.Left.Walk() + e.Right.Walk()
}

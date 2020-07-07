package ast

type MulExpr struct {
	Left  Ast
	Right Ast
}

func NewMulExpr(left, right Ast) *MulExpr {
	return &MulExpr{
		Left:  left,
		Right: right,
	}
}

func (e *MulExpr) Walk() int64 {
	return e.Left.Walk() * e.Right.Walk()
}

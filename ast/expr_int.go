package ast

type IntExpr struct {
	Val int64
}

func NewIntExpr(val int64) *IntExpr {
	return &IntExpr{
		Val: val,
	}
}

func (e *IntExpr) Walk() int64 {
	return e.Val
}

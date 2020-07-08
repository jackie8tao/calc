package ast

type (
	// IntExpr
	IntExpr struct {
		Val int64
	}

	// AddExpr
	AddExpr struct {
		Left  Ast
		Right Ast
	}

	// SubExpr
	SubExpr struct {
		Left  Ast
		Right Ast
	}

	// DivExpr
	DivExpr struct {
		Left  Ast
		Right Ast
	}

	// MulExpr
	MulExpr struct {
		Left  Ast
		Right Ast
	}
)

func NewAddExpr(left, right Ast) *AddExpr {
	return &AddExpr{
		Left:  left,
		Right: right,
	}
}

func (e *AddExpr) Walk() int64 {
	return e.Left.Walk() + e.Right.Walk()
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

func NewIntExpr(val int64) *IntExpr {
	return &IntExpr{
		Val: val,
	}
}

func (e *IntExpr) Walk() int64 {
	return e.Val
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

func NewSubExpr(left, right Ast) *SubExpr {
	return &SubExpr{
		Left:  left,
		Right: right,
	}
}

func (e *SubExpr) Walk() int64 {
	return e.Left.Walk() - e.Right.Walk()
}

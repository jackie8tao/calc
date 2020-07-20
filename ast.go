package main

// operators
const (
	OpAdd Operator = iota + 1
	OpSub
	OpMul
	OpDiv
	OpInt
)

// Op expression operator
type Operator int

// Ast abstract syntax tree
type Ast interface {
	// Op return ast operator
	Op() Operator
}

/* ------------------------------------------------------ */
// IntAst integer number
type IntAst struct {
	Num int64 // lexical number
}

func (t *IntAst) Op() Operator {
	return OpInt
}

/* ------------------------------------------------------ */
// AddAst add expression
type AddAst struct {
	Left  Ast
	Right Ast
}

func (t *AddAst) Op() Operator {
	return OpAdd
}

/* ------------------------------------------------------ */
// SubAst sub expression
type SubAst struct {
	Left  Ast
	Right Ast
}

func (t *SubAst) Op() Operator {
	return OpSub
}

/* ------------------------------------------------------ */
// MulAst multiple expression
type MulAst struct {
	Left  Ast
	Right Ast
}

func (t *MulAst) Op() Operator {
	return OpMul
}

/* ------------------------------------------------------ */
// DivAst divide expression
type DivAst struct {
	Left  Ast
	Right Ast
}

func (t *DivAst) Op() Operator {
	return OpDiv
}

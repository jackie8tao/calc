package main

// Visit visit the tree and get result.
func Visit(ast Ast) (ret int64) {
	switch ast.Op() {
	case OpInt:
		v, ok := ast.(*IntAst)
		if !ok {
			panic(ErrInvalidAst)
		}
		ret = v.Num
	case OpAdd:
		v, ok := ast.(*AddAst)
		if !ok {
			panic(ErrInvalidAst)
		}
		ret = Visit(v.Left) + Visit(v.Right)
	case OpSub:
		v, ok := ast.(*SubAst)
		if !ok {
			panic(ErrInvalidAst)
		}
		ret = Visit(v.Left) - Visit(v.Right)
	case OpMul:
		v, ok := ast.(*MulAst)
		if !ok {
			panic(ErrInvalidAst)
		}
		ret = Visit(v.Left) * Visit(v.Right)
	case OpDiv:
		v, ok := ast.(*DivAst)
		if !ok {
			panic(ErrInvalidAst)
		}
		ret = Visit(v.Left) / Visit(v.Right)
	default:
		panic(ErrInvalidAst)
	}
	return
}

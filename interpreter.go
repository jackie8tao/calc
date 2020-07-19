package main

// Visitor walk through the ast, calculate the expression.
type Visitor interface {
	// Visit walk through the ast and calculate.
	Visit(ast Ast) int64
}

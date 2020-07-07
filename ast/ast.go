package ast

type Ast interface {
	Walk() int64
}

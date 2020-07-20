package main

import "errors"

var (
	ErrIllegalToken    = errors.New("illegal token")                // ErrIllegalToken illegal token
	ErrEOZ             = errors.New("end of stream")                // ErrEOZ end of stream
	ErrEmpty           = errors.New("empty stream")                 // ErrEmpty empty stream
	ErrUnexpectedToken = errors.New("unexpected token")             // ErrUnexpectedToken unexpected token
	ErrInvalidAst      = errors.New("invalid abstract syntax tree") // ErrInvalidAst invalid ast
	ErrZeroOperand     = errors.New("zero operand used in divide")  // ErrZeroOperand zero operand
)

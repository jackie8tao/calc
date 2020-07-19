package main

import "errors"

var (
	// ErrIllegalToken illegal token
	ErrIllegalToken = errors.New("illegal token")

	// ErrEOZ end of stream
	ErrEOZ = errors.New("end of stream")

	// ErrEmpty empty stream
	ErrEmpty = errors.New("empty stream")

	// ErrUnexpectedToken unexpected token
	ErrUnexpectedToken = errors.New("unexpected token")
)

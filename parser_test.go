package main

import "testing"

func TestParser_Next(t *testing.T) {
	p := NewParser("1+2+3")
	a := p.Next()
	if a == nil {
		t.Fail()
	}
}

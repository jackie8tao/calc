package main

import "testing"

func TestNew(t *testing.T) {
	p := New("1+2+3")
	if p == nil {
		t.Fail()
	}
}

func TestParser_Next(t *testing.T) {
	p := New("1+2+3")
	a := p.Next()
	if a == nil {
		t.Fail()
		return
	}
	if a.Walk() != 6 {
		t.Fail()
	}
}

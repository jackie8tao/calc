package main

import "testing"

func TestLexer_Get(t *testing.T) {
	lex := NewLexer("12 + 5 * (8 + 9) - 6 / 2")
	for {
		tok, val := lex.Get()
		if lex.Err() != nil {
			if lex.Err() == ErrEOZ {
				return
			}
			t.Fail()
		}
		t.Logf("token: %d, value: %s", tok, val)
	}
}

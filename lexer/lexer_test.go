package lexer

import "testing"

func TestLexer_Get(t *testing.T) {
	l := New("12+5-3")
	for {
		tok, val := l.Get()
		if l.err != nil {
			t.Log(l.err)
			return
		}
		t.Logf("tok: %d, val: %s", tok, val)
	}
}

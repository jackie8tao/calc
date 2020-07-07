package lexer

import (
	"errors"

	"calc/token"
)

const EOF = rune(-1)

type Lexer struct {
	err    error
	ch     rune
	offset int
	src    []rune
	buf    []rune
}

func New(src string) *Lexer {
	rsrc := []rune(src)
	return &Lexer{
		src:    rsrc,
		offset: 0,
		ch:     rsrc[0],
		buf:    []rune{},
	}
}

func isEOF(c rune) bool {
	return c == EOF
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func isWhitespace(c rune) bool {
	return c == ' ' || c == '\t'
}

func (l *Lexer) save() {
	l.buf = append(l.buf, l.ch)
}

func (l *Lexer) next() {
	l.offset++
	if l.offset >= len(l.src) {
		l.ch = EOF
		return
	}
	l.ch = l.src[l.offset]
}

func (l *Lexer) nextAndSave() {
	l.save()
	l.next()
}

func (l *Lexer) reset() {
	l.buf = make([]rune, 0)
}

func (l *Lexer) Get() (tok token.Token, val string) {
	l.reset()
	for {
		switch c := l.ch; {
		case isEOF(c):
			l.err = errors.New("end of the stream")
			return
		case isWhitespace(c):
			l.next()
		case isDigit(c):
			l.nextAndSave()
			for isDigit(l.ch) {
				l.nextAndSave()
			}
			tok = token.INT
			val = string(l.buf)
			return
		default:
			tok = token.LookupOperator(c)
			if tok == token.ILLEGAL {
				l.err = errors.New("illegal token")
				return
			}
			l.nextAndSave()
			val = string(l.buf)
			return
		}
	}
}

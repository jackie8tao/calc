package main

import (
	"calc/token"
)

// EOZ end of file
const EOZ = rune(-1)

// Lexer parse token from input stream
type Lexer struct {
	err    error
	ch     rune
	offset int
	src    []rune
	buf    []rune
}

// NewLexer create lexer upon source stream
func NewLexer(src string) *Lexer {
	val := []rune(src)
	if len(val) <= 0 {
		panic(ErrEmpty)
	}

	return &Lexer{
		ch:     val[0],
		src:    val,
		buf:    make([]rune, 0),
		offset: 0,
	}
}

func isEOF(c rune) bool {
	return c == EOZ
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
		l.ch = EOZ
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

// Err return lexer error
func (l *Lexer) Err() error {
	return l.err
}

// Get return next token from stream
func (l *Lexer) Get() (tok token.Token, val string) {
	l.reset()

	for {
		switch c := l.ch; {
		case isEOF(c):
			l.err = ErrEOZ
			return
		case isWhitespace(c):
			l.next()
		case isDigit(c):
			l.nextAndSave()
			for isDigit(l.ch) {
				l.nextAndSave()
			}
			tok = token.INT
			goto success
		default:
			tok = token.LookupOperator(c)
			l.nextAndSave()
			goto success
		}
	}
success:
	val = string(l.buf)
	return
}

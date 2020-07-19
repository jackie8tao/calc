package main

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

// check whether current char is EOZ.
func (l *Lexer) isEOZ() bool {
	return l.ch == EOZ
}

// check whether current char is digit.
func (l *Lexer) isDigit() bool {
	return '0' <= l.ch && l.ch <= '9'
}

// check whether current char is whitespace.
func (l *Lexer) isWhitespace() bool {
	return l.ch == ' ' || l.ch == '\t'
}

// save current char into buffer.
func (l *Lexer) save() {
	l.buf = append(l.buf, l.ch)
}

// eat current char and read next.
func (l *Lexer) next() {
	l.offset++
	if l.offset >= len(l.src) {
		l.ch = EOZ
		return
	}
	l.ch = l.src[l.offset]
}

// save current char and read next.
func (l *Lexer) nextAndSave() {
	l.save()
	l.next()
}

// clear buffer.
func (l *Lexer) reset() {
	l.buf = make([]rune, 0)
}

// Err return lexer error
func (l *Lexer) Err() error {
	return l.err
}

// Get return next token from stream
func (l *Lexer) Get() (tok Token, val string) {
	l.reset()

	for {
		switch {
		case l.isEOZ():
			l.err = ErrEOZ
			return
		case l.isWhitespace():
			l.next()
		case l.isDigit():
			l.nextAndSave()
			for l.isDigit() {
				l.nextAndSave()
			}
			tok = TokInt
			goto success
		default:
			tok = lookupToken(l.ch)
			l.nextAndSave()
			goto success
		}
	}
success:
	val = string(l.buf)
	return
}

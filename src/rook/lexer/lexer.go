// lexer.go

package lexer

import (
	"rook/token"
)

// Lexer
// makes tokens
type Lexer struct {
	input        string
	position     int // points to current char
	readPosition int // points to after current char
	ch           byte
	Tokens       []token.Token
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	// we call readChar here to put the Lexer in it's default state.
	l.readChar()
	return l
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) IsAtEnd() bool {
	if l.readPosition >= len(l.input) {
		return true
	}
	return false
}

func (l *Lexer) Tokenize() []token.Token {
	for !l.IsAtEnd() {
		l.Tokens = append(l.Tokens, l.NextToken())
	}

	return l.Tokens
}

// creates a new token.
func (l *Lexer) newToken(tokenType token.Type, lit string) token.Token {
	return token.Token{tokenType, lit, l.position, len(lit)}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = l.newToken(token.EQ, literal)
		} else {
			tok = l.newToken(token.ASSIGN, string(l.ch))
		}
	case '!':
		tok = l.newToken(token.BANG, string(l.ch))
	case '+':
		tok = l.newToken(token.PLUS, string(l.ch))
	case '-':
		tok = l.newToken(token.MINUS, string(l.ch))
	case '*':
		tok = l.newToken(token.ASTERISK, string(l.ch))
	case '/':
		tok = l.newToken(token.SLASH, string(l.ch))
	default:
		if isLetter(l.ch) {
			tok.Literal = l.identifier()
			tok.Type = token.Lookup(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) identifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// advances the position forward, or doesn't,
// if we're at the end
func (l *Lexer) readChar() {
	// if we're past the end, return 0
	if l.readPosition >= len(l.input) {
		l.ch = 0

		// else return an char
	} else {
		l.ch = l.input[l.readPosition]
	}
	// increment for the next iteration.
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

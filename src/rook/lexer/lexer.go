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
	l.stripTrailingWhitespace()
	l.ensureTerminalByte()
	l.readChar()
	return l
}

func (* Lexer) stripTrailingWhitespace() {
	str := []byte(l.input)
	last := len(str) -1
	lastCh := str[last]
}

func (l *Lexer) ensureTerminalByte() {
	str := []byte(l.input)
	last := len(str) - 1
	lastCh := str[last]

	if lastCh != 0 {
		str := []byte(l.input)
		str = append(str, 0)
		l.input = string(str)
	}
}

// should trim the trailing whitespace,
// or we should redesign this thing so that trailing whitespace doesn't trip it up.
func (l *Lexer) trimTrailingWhitespace() {

}

func (l *Lexer) peekChar() byte {
	return l.peek(0)
}

// peeks an arbitrary length.
func (l *Lexer) peek(position int) byte {
	pos := l.readPosition + position
	if pos >= len(l.input) {
		return 0
	} else {
		return l.input[pos]
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

	if l.Tokens[len(l.Tokens) - 1].Type != token.EOF {
		l.Tokens = append(l.Tokens, token.Token{Type: token.EOF, Literal: "EOF", Start: len(l.input), Length: len("EOF")})
	}

	return l.Tokens
}

// creates a new token.
func (l *Lexer) newToken(tokenType token.Type, lit string) token.Token {
	return token.Token{Type: tokenType, Literal: lit, Start: l.position, Length: len(lit)}
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
	case '0':
		if l.peekChar() == 'x' {
			tok.Type = token.BYTE
			tok.Literal = l.readByte()
			return tok
		} else {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.identifier()
			tok.Type = token.Lookup(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			// probably redundant now.
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			if isWhitespace(l.ch) && l.IsAtEnd() == true {
				l.Tokens = append(l.Tokens, token.Token{Type: token.EOF, Literal: "EOF", Start: len(l.input), Length: len("EOF")})
			}
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

func isHex(ch byte) bool {
	return '0' <= ch && ch <= '9' || 'A' <= ch && ch <= 'F' || 'a' <= ch && ch <= 'f'
}

// Lexer.readChar()
// Advances the position forward, or doesn't, if we're at the end.
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

func (l *Lexer) readByte() string {
	position := l.position
	if isHex(l.peek(1)) && isHex(l.peek(2)) {
		l.readChar()
		l.readChar()
		l.readChar()
		l.readChar()
	} else {
		// error lexing the Hexadecimal number.
		// TODO: Add an error token here for parsing reading the number wrong.
	}

	return l.input[position:l.position]
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
	for isWhitespace(l.ch) {
		l.readChar()
	}
}

func isWhitespace(ch byte) bool {
	return (ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r')
}

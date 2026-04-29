// lexer_test.go

package lexer

import (
	"rook/token"
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestLexerNew(t *testing.T) {
	l := New("5 whatever 5")
	tokens := l.Tokenize()
	assert.Equal(t, 4, len(tokens), "Not enough tokens.")

	assert.Equal(t, token.Type(token.INT), tokens[0].Type, "Expecting INT Token.")
	assert.Equal(t, token.Type(token.IDENT), tokens[1].Type, "Expecting IDENT Token.")
	assert.Equal(t, token.Type(token.INT), tokens[2].Type, "Expecting INT Token.")
	assert.Equal(t, token.Type(token.EOF), tokens[3].Type, "Expecting EOF Token.")
}

func TestLexing(t *testing.T) {
	l := New(`5 + 5
999 - 99
50 * 50
60 / 6
hello friend
!true
whatever = 10
whatever == 10
const thing = 5
0xFF    `)

	tokens := l.Tokenize()
	assert.Equal(t, 28, len(tokens), "Not enough tokens.")
	assert.Equal(t, token.Type(token.PLUS), tokens[1].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.MINUS), tokens[4].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.ASTERISK), tokens[7].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.SLASH), tokens[10].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.BANG), tokens[14].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.ASSIGN), tokens[17].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.EQ), tokens[20].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.CONST), tokens[22].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.IDENT), tokens[23].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.ASSIGN), tokens[24].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.INT), tokens[25].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.BYTE), tokens[26].Type, "Token not of the right type.")
	assert.Equal(t, "EOF", tokens[27].Literal, "Not the right literal")
	assert.Equal(t, token.Type(token.EOF), tokens[27].Type, "Expecting EOF Token.")
}

func TestGroupers(t *testing.T) {
	l := New(`
(){}[]
`)
	tokens := l.Tokenize()
	assert.Equal(t, 7, len(tokens), "Not enough tokens.")
	assert.Equal(t, token.Type(token.LPAREN), tokens[0].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.RPAREN), tokens[1].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.LBRACE), tokens[2].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.RBRACE), tokens[3].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.LBRACKET), tokens[4].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.RBRACKET), tokens[5].Type, "Token not of the right type.")
	assert.Equal(t, "EOF", tokens[6].Literal, "Not the right literal")
	assert.Equal(t, token.Type(token.EOF), tokens[6].Type, "Expecting EOF Token.")
}


func TestDoubles(t *testing.T) {
	l := New(`
! != = == < > & && | ||
`)
	tokens := l.Tokenize()
	assert.Equal(t, 11, len(tokens), "Not enough tokens.")
	assert.Equal(t, token.Type(token.BANG), tokens[0].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.NOT_EQ), tokens[1].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.ASSIGN), tokens[2].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.EQ), tokens[3].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.LT), tokens[4].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.GT), tokens[5].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.AMPERSAND), tokens[6].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.AND), tokens[7].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.PIPE), tokens[8].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.OR), tokens[9].Type, "Token not of the right type.")
	assert.Equal(t, "EOF", tokens[10].Literal, "Not the right literal")
}

func TestKeywords(t *testing.T) {
	l := New(`
fun   pub    type   let
const true   false  if
else  for    switch case
return
`)

	tokens := l.Tokenize()
	assert.Equal(t, 14, len(tokens), "Not enough tokens.")

	assert.Equal(t, token.Type(token.FUNCTION), tokens[0].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.PUBLIC), tokens[1].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.TYPE), tokens[2].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.LET), tokens[3].Type, "Token not of the right type.")

	assert.Equal(t, token.Type(token.CONST), tokens[4].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.TRUE), tokens[5].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.FALSE), tokens[6].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.IF), tokens[7].Type, "Token not of the right type.")

	assert.Equal(t, token.Type(token.ELSE), tokens[8].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.FOR), tokens[9].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.SWITCH), tokens[10].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.CASE), tokens[11].Type, "Token not of the right type.")

	assert.Equal(t, token.Type(token.RETURN), tokens[12].Type, "Token not of the right type.")

	assert.Equal(t, token.Type(token.EOF), tokens[13].Type, "Token not of the right type.")

}

// // Keywords
// 	FUNCTION = "FUNCTION"
// 	PUBLIC   = "PUBLIC"
// 	TYPE     = "TYPE"
// 	LET      = "LET"
// 	CONST    = "CONST"
// 	TRUE     = "TRUE"
// 	FALSE    = "FALSE"
// 	IF       = "IF"
// 	ELSE     = "ELSE"
// 	FOR      = "FOR"
// 	SWITCH   = "SWITCH"
// 	CASE     = "CASE"
// 	RETURN   = "RETURN"


func TestNullTerminatingByte(t *testing.T) {
	l := New(`5 + 5`)
	tokens := l.Tokenize()

	if 4 != len(tokens) {
		assert.Equal(t, 4, len(tokens), "Not enough tokens.")
		for _, tk := range tokens {
			fmt.Printf("%s", tk.String())
		}
	}

	assert.Equal(t, "EOF", tokens[3].Literal, "Not the right literal")
	assert.Equal(t, token.Type(token.EOF), tokens[3].Type, "Expecting EOF Token.")
}

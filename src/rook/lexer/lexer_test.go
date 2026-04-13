// lexer_test.go

package lexer

import (
	"rook/token"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexerNew(t *testing.T) {
	l := New("5 whatever 5")
	tokens := l.Tokenize()
	assert.Equal(t, 3, len(tokens), "Not enough tokens.")

	assert.Equal(t, token.Type(token.INT), tokens[0].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.IDENT), tokens[1].Type, "Token not of the right type.")
	assert.Equal(t, token.Type(token.INT), tokens[2].Type, "Token not of the right type.")
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
const thing = 5`)

	tokens := l.Tokenize()
	assert.Equal(t, 26, len(tokens), "Not enough tokens.")

	assert.Equal(t,
		token.Type(token.PLUS), tokens[1].Type,
		"Token not of the right type.")

	assert.Equal(t,
		token.Type(token.MINUS), tokens[4].Type,
		"Token not of the right type.")

	assert.Equal(t,
		token.Type(token.ASTERISK), tokens[7].Type,
		"Token not of the right type.")

	assert.Equal(t,
		token.Type(token.SLASH), tokens[10].Type,
		"Token not of the right type.")

	assert.Equal(t,
		token.Type(token.BANG), tokens[14].Type,
		"Token not of the right type.")

	assert.Equal(t,
		token.Type(token.ASSIGN), tokens[17].Type,
		"Token not of the right type.")

	assert.Equal(t,
		token.Type(token.EQ), tokens[20].Type,
		"Token not of the right type.")

	assert.Equal(t,
		token.Type(token.CONST), tokens[22].Type,
		"Token not of the right type.")

}

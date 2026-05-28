// ast_test.go

package ast

import (
	"rook/lexer"
	"rook/token"
	"testing"
	// "fmt"
	"github.com/stretchr/testify/assert"
)

func TestMakeNodes(t *testing.T) {
	l := lexer.New("mod whatever")
	tokens := l.Tokenize()
	assert.Equal(t, 4, len(tokens), "Not enough tokens.")

	a := Module(tokens[0])
	b := Identifier(tokens[1])

	assert.Equal(t, token.MODULE, string(a.Type()), "Expecting this node to be a module token.")
	assert.Equal(t, token.IDENT, string(b.Type()), "Expecting IDENT Token.")

}

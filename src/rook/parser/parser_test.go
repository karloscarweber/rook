// parser_test.go

package parser

import (
	"testing"
	// "fmt"
	"github.com/stretchr/testify/assert"
)

func TestParserNew(t *testing.T) {

	p := New("mod rook")
	p.Parse()

	// tokens := l.Tokenize()
	if len(p.Nodes) != 1 {
		p.DebugErrors()
		p.DebugTokens()
		assert.Equal(t, 1, len(p.Nodes), "Not enough Nodes")
	}

	// assert.Equal(t, token.Type(token.INT), tokens[0].Type, "Expecting INT Token.")
	// assert.Equal(t, token.Type(token.IDENT), tokens[1].Type, "Expecting IDENT Token.")
	// assert.Equal(t, token.Type(token.INT), tokens[2].Type, "Expecting INT Token.")
	// assert.Equal(t, token.Type(token.EOF), tokens[3].Type, "Expecting EOF Token.")
}

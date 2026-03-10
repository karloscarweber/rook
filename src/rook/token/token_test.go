package token

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.True(t, true, "True is True")
}

func TestKeywordLookup(t *testing.T) {
	assert.Equal(t, TokenType(FUNCTION), Lookup("fun"),
		"Expecting 'fun' Token.")
	assert.Equal(t, TokenType(PUBLIC), Lookup("pub"),
		"Expecting 'pub' Token.")
	assert.Equal(t, TokenType(TYPE), Lookup("type"),
		"Expecting 'type' Token.")
	assert.Equal(t, TokenType(LET), Lookup("let"),
		"Expecting 'let' Token.")
	assert.Equal(t, TokenType(CONST), Lookup("const"),
		"Expecting 'const' Token.")
	assert.Equal(t, TokenType(TRUE), Lookup("true"),
		"Expecting 'true' Token.")
	assert.Equal(t, TokenType(FALSE), Lookup("false"),
		"Expecting 'false' Token.")
	assert.Equal(t, TokenType(IF), Lookup("if"),
		"Expecting 'if' Token.")
	assert.Equal(t, TokenType(ELSE), Lookup("else"),
		"Expecting 'else' Token.")
	assert.Equal(t, TokenType(FOR), Lookup("for"),
		"Expecting 'for' Token.")
	assert.Equal(t, TokenType(SWITCH), Lookup("switch"),
		"Expecting 'switch' Token.")
	assert.Equal(t, TokenType(CASE), Lookup("case"),
		"Expecting 'case' Token.")
	assert.Equal(t, TokenType(RETURN), Lookup("return"),
		"Expecting 'return' Token.")
}

func TestIdentifierLookup(t *testing.T) {
	assert.Equal(t, TokenType(IDENT), Lookup("foo"),
		"Expecting 'foo' Token.")
	assert.Equal(t, TokenType(IDENT), Lookup("bar"),
		"Expecting 'bar' Token.")
	assert.Equal(t, TokenType(IDENT), Lookup("baz"),
		"Expecting 'baz' Token.")
}

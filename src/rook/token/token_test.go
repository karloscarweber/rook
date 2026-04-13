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
	assert.Equal(t, Type(FUNCTION), Lookup("fun"),
		"Expecting 'fun' Token.")
	assert.Equal(t, Type(PUBLIC), Lookup("pub"),
		"Expecting 'pub' Token.")
	assert.Equal(t, Type(TYPE), Lookup("type"),
		"Expecting 'type' Token.")
	assert.Equal(t, Type(LET), Lookup("let"),
		"Expecting 'let' Token.")
	assert.Equal(t, Type(CONST), Lookup("const"),
		"Expecting 'const' Token.")
	assert.Equal(t, Type(TRUE), Lookup("true"),
		"Expecting 'true' Token.")
	assert.Equal(t, Type(FALSE), Lookup("false"),
		"Expecting 'false' Token.")
	assert.Equal(t, Type(IF), Lookup("if"),
		"Expecting 'if' Token.")
	assert.Equal(t, Type(ELSE), Lookup("else"),
		"Expecting 'else' Token.")
	assert.Equal(t, Type(FOR), Lookup("for"),
		"Expecting 'for' Token.")
	assert.Equal(t, Type(SWITCH), Lookup("switch"),
		"Expecting 'switch' Token.")
	assert.Equal(t, Type(CASE), Lookup("case"),
		"Expecting 'case' Token.")
	assert.Equal(t, Type(RETURN), Lookup("return"),
		"Expecting 'return' Token.")
}

func TestIdentifierLookup(t *testing.T) {
	assert.Equal(t, Type(IDENT), Lookup("foo"),
		"Expecting 'foo' Token.")
	assert.Equal(t, Type(IDENT), Lookup("bar"),
		"Expecting 'bar' Token.")
	assert.Equal(t, Type(IDENT), Lookup("baz"),
		"Expecting 'baz' Token.")
}
